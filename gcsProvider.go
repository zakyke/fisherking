package fisherking

import (
	"context"
	"io"
	"strings"

	"golang.org/x/oauth2/google"
	gcsstorage "google.golang.org/api/storage/v1"
)

type GCS struct {
}

func (GCS) GetWithContect(contect context.Context, path string) FileGetter {
	//Listern to cancel channel.
	return GCS{}.Get
}
func (GCS) Get(path string) (io.Reader, error) {
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

	GCSObject, err := service.Objects.Get(bucket, object).Fields("MediaLink").Do()
	if err != nil {
		return nil, err
	}
	rsp, err := gcsc.Get(GCSObject.MediaLink)
	if err != nil {
		return nil, err
	}
	return rsp.Body, nil
}

func parseGCSBucket(path string) (bucket, file string) {
	bs := strings.Index(path, gcsPrefix)
	be := strings.Index(path[bs:], pathSeperator)
	return path[bs : be-1], path[be:]
}
