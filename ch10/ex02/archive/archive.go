package archive

import (
	"io"
)

type File struct {
	Name string
	Body io.ReadCloser
}

type Archive interface {
	Next() (*File, error)
}
