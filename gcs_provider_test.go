package fisherking

import (
	"io"
	"os"
	"testing"
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

func TestGCSGetReader(t *testing.T) {
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

// func TestGCSGetReader(t *testing.T) {
// 	//Create bucket
// 	//Create file
// 	//Read the file
// }
