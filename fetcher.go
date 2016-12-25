//Package fisherking ...
/*
*
*File fetcher get return an io.Reader with data from a file in the web or local disk.
*Features:
*   Resolve the file provider from the path GCS S3 FS
*   Its possible to add more providers by implementing Fetcher interface
*   Save activity log.
*
*   Accept Contect for cancelation and timeouts
*   Clean automatically after a while
*   Pass credencial in context
*
*
 */
package fisherking

import (
	"context"
	"io"
	"strings"
)

//Provider  represent a provider like GCS, S3, ..
type Provider struct {
	Name string
	context.Context
}

const (
	gcsPrefix     = `gs://`
	s3Prefix      = `s3://`
	fsPrefix      = `file://`
	htmPrefix     = `html://`
	pathSeperator = `/`
)

//FileGetter ...
type FileGetter func(string) (io.Reader, error)
type FilePutter func(string) (io.Writer, error)
type FileDeleter func(string, string) error

//Fetcher an interface a provider should implement
type Fetcher interface {
	GetWithContext(contect context.Context, source string) FileGetter
	Get(source string) (io.Reader, error)
}

type Putter interface {
	PutWithContext(context context.Context, source, destination string) FilePutter
	Put(destination string) (io.Writer, error)
}

//GetWithContext can use for multiple files.
func GetWithContext(contect context.Context, source string) FileGetter {
	return Get
}

//Get a single file
func Get(path string) (io.Reader, error) {
	p := providerFactory(path)
	return p.Get(path)
}

func providerFactory(path string) Fetcher {
	ind := strings.Index(path, `://`)
	indicator := path[:ind+3]
	switch indicator {
	case fsPrefix:
		return fs{}
	}
	return nil
}
