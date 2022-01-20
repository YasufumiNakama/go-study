package tar

import (
	"archive/tar"
	"io"
	"io/ioutil"

	"ex02/archive" //"github.com/kdama/gopl/ch10/ex02/archive"
)

const magicNumber = "ustar" // https://ja.wikipedia.org/wiki/Tar
const offset = 257

func init() {
	archive.RegisterFormat("tar", magicNumber, offset, ReadArchive)
}

type Archive struct {
	tr *tar.Reader
}

func (a *Archive) Next() (*archive.File, error) {
	h, err := a.tr.Next()
	if err == io.EOF {
		return nil, io.EOF
	}
	if err != nil {
		return nil, err
	}
	return &archive.File{
		Name: h.Name,
		Body: ioutil.NopCloser(a.tr),
	}, nil
}

func ReadArchive(r io.Reader) (archive.Archive, error) {
	tr := tar.NewReader(r)
	return &Archive{tr: tr}, nil
}
