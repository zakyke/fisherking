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

import "io"
import "context"

type Provider struct {
	Name string
	context.Context
}

const (
	gcsPrefix     = `gs://`
	s3Prefix      = `s3://`
	fsPrefix      = `file://`
	pathSeperator = `/`
)

type FileGetter func(path string) (io.Reader, error)

type Fetcher interface {
	GetWithContect(contect context.Context, path string) FileGetter
	Get(path string) (io.Reader, error)
}

type Fetch struct {
}

func (Fetch) GetWithContect(contect context.Context, path string) FileGetter {
	return nil
}

func (Fetch) Get(path string) (io.Reader, error) {
	return nil, nil
}

func ProviderFactory(path string) Fetcher {
	//g,f,s
	if path[0] == 'g' {
		return GCS{}
	}
	return nil
}
