package fisherking

import (
	"context"
	"io"
	"log"
	"net/http"
	"strings"

	"bytes"

	"golang.org/x/oauth2/google"
	gcsstorage "google.golang.org/api/storage/v1"
)

type gcs struct {
	context.Context
}

func (gcs) Get(path string) (io.Reader, error) {
	service, client, err := newGcsService()
	if err != nil {
		return nil, err
	}
	//TODO cancel the call if context cancel or timeout
	bucket, object := parseGCSBucket(path)

	GCSObject, err := service.Objects.Get(bucket, object).Fields(`mediaLink`).Do()
	if err != nil {
		return nil, err
	}
	rsp, err := client.Get(GCSObject.MediaLink)
	if err != nil {
		return nil, err
	}
	defer checkClose(rsp.Body)

	//Drain the body in order to reuse the http connection
	data := bytes.NewBuffer(nil)
	_, e := io.Copy(data, rsp.Body)
	if e != nil {
		return nil, e
	}
	rtn := bytes.NewReader(data.Bytes())
	return rtn, nil
}

func parseGCSBucket(path string) (bucket, file string) {
	gslen := len(gcsPrefix)
	be := strings.Index(path[gslen:], linPathSeperator)
	return path[gslen : be+gslen], path[be+gslen+1:]
}

func newGcsService() (*gcsstorage.Service, *http.Client, error) {
	gcsc, err := google.DefaultClient(context.Background(), gcsstorage.DevstorageFullControlScope)
	if err != nil {
		return nil, nil, err
	}
	service, err := gcsstorage.New(gcsc)
	if err != nil {
		return nil, nil, err
	}
	return service, gcsc, nil
}

func (g gcs) Put(destination string, data io.Reader) error {
	service, _, err := newGcsService()
	if err != nil {
		return err
	}
	bucketName, fileName := parseGCSBucket(destination)
	object := &gcsstorage.Object{Name: fileName}
	metadata := g.Value(`metadata`)
	if metadata != nil {
		if md, ok := metadata.(map[string]string); ok {
			object.Metadata = md
		}
	}

	if _, err := service.Objects.Insert(bucketName, object).Media(data).Do(); err != nil {
		log.Printf("failed to upload file to GCS %s: %v\n", fileName, err)
		return err
	}
	return nil
}
