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
	"errors"
	"io"
	"strings"
)

//Provider  represent a provider like GCS, S3, ..
type Provider struct {
	Name string
	context.Context
}

const (
	gcsPrefix         = `gs://`
	s3Prefix          = `s3://`
	fsPrefix          = `file://`
	httpPrefix        = `http://`
	httpsPrefix       = `https://`
	pathSeperator     = `/`
	providerDelimiter = `://`
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

type Deleter interface {
	DeleteWithContext(context context.Context, source, destination string) FilePutter
	Delete(destination string) (io.Writer, error)
}

//GetWithContext can use for multiple files.
func GetWithContext(cxt context.Context, path string) FileGetter {
	p := providerFactory(cxt, path)
	if p == nil {
		return nil
	}
	return p.Get
}

//Get a single file
func Get(path string) (io.Reader, error) {
	p := providerFactory(nil, path)
	if p == nil {
		return nil, errors.New(`invalid provider prefix`)
	}
	return p.Get(path)
}

func providerFactory(ctx context.Context, path string) Fetcher {
	ind := strings.Index(path, providerDelimiter)
	indicator := path[:ind+3]
	switch indicator {
	case fsPrefix:
		return fs{ctx}
	case s3Prefix:
		return s3{ctx}
	case gcsPrefix:
		return gcs{ctx}
	case httpPrefix, httpsPrefix:
		return http{ctx}
	}
	return nil
}
