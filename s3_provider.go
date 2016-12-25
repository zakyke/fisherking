package fisherking

import (
	"context"
	"io"
	"strings"

	"golang.org/x/oauth2/google"
	gcsstorage "google.golang.org/api/storage/v1"
)

type s3 struct {
}

func (s3) GetWithContect(contect context.Context, path string) FileGetter {
	//Listern to cancel channel.
	return gcs{}.Get
}
func (s3) Get(path string) (io.Reader, error) {
	bucket, object := parseGCSBucket(path)
	gcsc, err := google.DefaultClient(context.Background(), gcsstorage.DevstorageFullControlScope)
	if err != nil {
		return nil, err
	}
	service, err := gcsstorage.New(gcsc)
	if err != nil {
		return nil, err
	}

	//TODO cancel the call if context cancel or timeout
	GCSObject, err := service.Objects.Get(bucket, object).Fields(`mediaLink`).Do()
	if err != nil {
		return nil, err
	}
	rsp, err := gcsc.Get(GCSObject.MediaLink)
	if err != nil {
		return nil, err
	}

	return rsp.Body, nil
}

func parseS3Bucket(path string) (bucket, file string) {
	be := strings.Index(path[4:], pathSeperator)
	return path[4 : be+4], path[be+4+1:]
}
