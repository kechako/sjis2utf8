package transform

import (
	"io"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

type Reader struct {
	reader *transform.Reader
}

func NewReader(rawReader io.Reader) Reader {
	reader := Reader{}
	reader.reader = transform.NewReader(rawReader, japanese.ShiftJIS.NewDecoder())

	return reader
}

func (r Reader) Read(p []byte) (int, error) {
	return r.reader.Read(p)
}
