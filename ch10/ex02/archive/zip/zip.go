package zip

import (
	"archive/zip"
	"bytes"
	"io"
	"io/ioutil"

	"ex02/archive" //"github.com/kdama/gopl/ch10/ex02/archive"
)

const magicNumber = "PK\003\004" // https://ja.wikipedia.org/wiki/ZIP_(ファイルフォーマット)
const offset = 0

func init() {
	archive.RegisterFormat("zip", magicNumber, offset, ReadArchive)
}

type Archive struct {
	zr      *zip.Reader
	current int
}

func (a *Archive) Next() (*archive.File, error) {
	if a.current >= len(a.zr.File) {
		return nil, io.EOF
	}
	f := a.zr.File[a.current]
	body, err := f.Open()
	if err != nil {
		return nil, err
	}
	return &archive.File{
		Name: f.Name,
		Body: body,
	}, nil
}

func ReadArchive(r io.Reader) (archive.Archive, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	readerAt := bytes.NewReader(b)
	zr, _ := zip.NewReader(readerAt, readerAt.Size())
	return &Archive{zr: zr}, nil
}
