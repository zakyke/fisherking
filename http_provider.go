package fisherking

import (
	"context"
	"io"
)

type http struct {
	context.Context
}

func (http) GetWithContext(contect context.Context, source string) FileGetter {
	panic("not implemented")
}

func (http) Get(source string) (io.Reader, error) {
	panic("not implemented")
}
func (http) PutWithContext(context context.Context, source string, destination string) FilePutter {
	panic("not implemented")
}

func (http) Put(destination string) (io.Writer, error) {
	panic("not implemented")
}
