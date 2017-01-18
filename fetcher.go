//Package fisherking implements access to a varaiaty of file storage.
//
//File fetcher get return an io.Reader with data from a file in the web or local disk.
//
//Features
//
//Resolve the file provider from the path GCS S3 FS
//
//Its possible to add more providers by implementing Fetcher interface
//
//Save activity log.
//
//Accept Contect for cancelation and timeouts
//
//Clean automatically after a while
//
//Use default credentials or pass credentials in context
package fisherking

import (
	"context"
	"errors"
	"io"
	"log"
	"reflect"
	"strings"
)

//Provider  represent a provider like GCS, S3, ..
type Provider struct {
	Name string
	context.Context
}

const (
	gcsPrefix        = `gs://`
	s3Prefix         = `s3://`
	fsPrefix         = `file://`
	httpPrefix       = `http://`
	httpsPrefix      = `https://`
	linPathSeperator = `/`
	//winPathSeperator  = `\`
	providerDelimiter = `://`
)

//FileGetter ...
// type FileGetter func(string) (io.Reader, error)
// type FilePutter func(context context.Context, destination string, data io.Writer) error

//type FileDeleter func(string, string) error
// type Deleter interface {
// 	DeleteWithContext(context context.Context, source, destination string) FilePutter
// 	Delete(destination string) (io.Writer, error)
// }

//Fisher an interface a provider should implement
type Fisher interface {
	Get(source string) (io.Reader, error)
	Put(destination string, data io.Reader) error
	//Delete(source string)  error
	//Move(source, destination string)  error
}

//Put write a Reader to destination without any context.
func Put(destination string, data io.Reader) error {
	return PutWithContext(context.Background(), destination, data)
}

//PutWithContext write a Reader to destination with context. metadata for example.
func PutWithContext(context context.Context, destination string, data io.Reader) error {
	p := providerFactory(context, destination)
	return p.Put(destination, data)
}

//GetWithContext can use for multiple files, cancelation...
func GetWithContext(cxt context.Context, source string) (io.Reader, error) {
	p := providerFactory(cxt, source)
	if p == nil {
		return nil, errors.New(`fail to parse path`)
	}
	return p.Get(source)
}

//Get a single file
func Get(source string) (io.Reader, error) {
	return GetWithContext(context.Background(), source)
}

func providerFactory(ctx context.Context, path string) Fisher {
	ind := strings.Index(path, providerDelimiter)
	indicator := path[:ind+len(providerDelimiter)]
	switch indicator {
	case fsPrefix:
		return fs{ctx}
	case s3Prefix:
		return s3{ctx}
	case gcsPrefix:
		return gcs{ctx}
	case httpPrefix, httpsPrefix:
		return htp{ctx}
	}
	return nil
}

func checkClose(ref io.Closer) {
	if e := ref.Close(); e != nil {
		log.Println(e, reflect.TypeOf(ref))
	}

}
