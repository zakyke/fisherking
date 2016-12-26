package fisherking

import (
	"context"
	"io"

	"os"
)

type fs struct {
	context.Context
}

func (f fs) GetWithContext(cxt context.Context, path string) FileGetter {
	return fs{cxt}.Get
}

func (f fs) Get(path string) (io.Reader, error) {
	path = path[6:]
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (f fs) PutWithContext(cxt context.Context, path string) FilePutter {
	return fs{cxt}.Put
}
func (f fs) Put(destination string) (io.Writer, error) {
	file, err := os.Open(destination)
	if err != nil {
		return nil, err
	}
	return file, nil
}
