// Derived from https://github.com/golang/go/blob/master/src/archive/format.go
package archive

import (
	"bufio"
	"errors"
	"io"
)

// ErrFormat indicates that archive encountered an unknown format.
var ErrFormat = errors.New("archive: unknown format")

// A format holds an archive format's name, magic header and how to decode it.
type format struct {
	name, magic string
	offset      int
	decode      func(io.Reader) (Archive, error)
}

// Formats is the list of registered formats.
var formats []format

// RegisterFormat registers an archive format for use by Decode.
// Name is the name of the format, like "zip" or "tar".
// Magic is the magic prefix that identifies the format's encoding. The magic
// string can contain "?" wildcards that each match any one byte.
// Decode is the function that decodes the encoded archive.
func RegisterFormat(name, magic string, offset int, decode func(io.Reader) (Archive, error)) {
	formats = append(formats, format{name, magic, offset, decode})
}

// A reader is an io.Reader that can also peek ahead.
type reader interface {
	io.Reader
	Peek(int) ([]byte, error)
}

// asReader converts an io.Reader to a reader.
func asReader(r io.Reader) reader {
	if rr, ok := r.(reader); ok {
		return rr
	}
	return bufio.NewReader(r)
}

// Match reports whether magic matches b. Magic may contain "?" wildcards.
func match(magic string, b []byte) bool {
	if len(magic) != len(b) {
		return false
	}
	for i, c := range b {
		if magic[i] != c && magic[i] != '?' {
			return false
		}
	}
	return true
}

// Sniff determines the format of r's data.
func sniff(r reader) format {
	for _, f := range formats {
		b, err := r.Peek(f.offset + len(f.magic))
		if err == nil && match(f.magic, b[f.offset:]) {
			return f
		}
	}
	return format{}
}

// Decode decodes an archive that has been encoded in a registered format.
// The string returned is the format name used during format registration.
// Format registration is typically done by an init function in the codec-
// specific package.
func ReadArchive(r io.Reader) (Archive, string, error) {
	rr := asReader(r)
	f := sniff(rr)
	if f.decode == nil {
		return nil, "", ErrFormat
	}
	m, err := f.decode(rr)
	return m, f.name, err
}
