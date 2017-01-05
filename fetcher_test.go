package fisherking

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestGCSRead(t *testing.T) {
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

func TestFSRead(t *testing.T) {
	r, e := Get(`file:///tmp/list.txt`)
	if e != nil {
		t.Log(e)
		t.Fail()
	}
	b := bytes.NewBufferString(``)
	_, err := io.Copy(b, r)
	t.Log(`error: `, err, b.String())
}

func TestFS2GCSCopy(t *testing.T) {
	_, e := Get(`file:///home/zaky/streamrail/go/src/github.com/streamrail/messaging/data/2016/01/01/10/56f91cd4d3e3660002000033/adsmanager
/http-collector-asia-multi-preemp-6xlj-adsmanager-1483228811571309144.gz`)
	if e != nil {
		t.Log(e)
		t.Fail()
	}

}
