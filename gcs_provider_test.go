package fisherking

import (
	"context"
	"io"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestGCSSplitPath(t *testing.T) {
	f := `gs://adsmanager_files/2016/12/11/770/j2016_12_11_770.jobs`
	b, o := parseGCSBucket(f)
	if b != `adsmanager_files` || o != `2016/12/11/770/j2016_12_11_770.jobs` {
		t.Log(b)
		t.Log(o)
		t.Fail()
	}
}

func TestReadGCS(t *testing.T) {
	f := `gs://adsmanager_files/2016/12/11/770/j2016_12_11_770.jobs`
	read, e := Get(f)
	if e != nil {
		t.Log(e)
		t.Fail()
	}
	if read == nil {
		t.Fail()
	}
	_, err := io.Copy(os.Stdout, read)
	t.Log(`error: `, err)
}

func TestReadWriteGCS(t *testing.T) {
	send := func(object string) {
		f := `file:///home/zaky/streamrail/go/src/github.com/streamrail/messaging/data/` + object
		r, e := Get(f)
		if e != nil {
			t.Log(e)
			t.Fail()
		}
		metadata := make(map[string]string)
		metadata[`rows`] = "230"
		metadata[`created`] = time.Now().Format(time.RFC3339)
		metadata[`bytes`] = "25555"

		ctx := context.WithValue(context.Background(), mdKey, metadata)
		e = PutWithContext(ctx, `gs://testtc/`+object, r)
		//	e = Put(`gs://testtc/`+object, r)
		if e != nil {
			t.Log(e)
			t.Fail()
		}
	}

	files := filesInFolder(`/home/zaky/streamrail/go/src/github.com/streamrail/messaging/data/2017/01/01/10/56f91cd4d3e3660002000033/adsmanager/`)

	for i := range files {
		send(files[i])
	}
}

func filesInFolder(path string) []string {
	files, _ := ioutil.ReadDir(path)
	trn := make([]string, len(files))
	for i, f := range files {
		trn[i] = `2017/01/01/10/56f91cd4d3e3660002000033/adsmanager/` + f.Name()
	}

	return trn
}

var mdKey interface{} = `metadata`

func TestReadWriteWithContextGCS(t *testing.T) {
	object := `2017/01/10/490/56f91cd4d3e3660002000033/adsmanager/http-collector-asia-multi-preemp-6lsv-adsmanager-1484035205409012184.gz`
	f := `gs://adsmanager_files/` + object
	r, e := Get(f)
	if e != nil {
		t.Log(e)
		t.Fail()
	}
	metadata := make(map[string]string)
	metadata[`rows`] = "230"
	metadata[`created`] = time.Now().Format(time.RFC3339)
	metadata[`bytes`] = "25555"

	ctx := context.WithValue(context.Background(), mdKey, metadata)
	e = PutWithContext(ctx, `gs://testtc/`+object, r)
	if e != nil {
		t.Log(e)
		t.Fail()
	}
}
