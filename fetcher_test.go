package fisherking

import (
	"bytes"
	"errors"
	"io"
	"os"
	"testing"
)

func ExampleGet() /*(io.Reader, error)*/ {
	f := `gs://bucket_name/2016/12/11/770/fileNmae.txt`
	/*read, err :=*/ _, _ = Get(f)
	//return read, err
	// Output: io.Reader and error
}

func ExamplePut() /*error*/ {
	send := func(object string) error {
		f := `file:///home/my/data/` + object
		r, e := Get(f)
		if e != nil {
			return e
		}
		e = Put(`gs://bucket_name/`+object, r)
		return e
	}

	files := filesInFolder(`/home/my/data`)

	for i := range files {
		if err := send(files[i]); err != nil {
			//return err
		}
	}
	//return nil
	// Output: Upload /home/my/data/* files to gs://bucket_name/ or an error.
}

func TestGCSRead(t *testing.T) {
	f := `gs://my_bucket/2016/12/11/770/j2016_12_11_770.jobs`
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
	_, e := Get(`file:///home/zaky/1483228811571309144.gz`)
	if e != nil {
		t.Log(e)
		t.Fail()
	}

}

func TestS32GCSCopy(t *testing.T) {
	r, e := Get(`s3://bucket/36_3p53hr4z.log`)
	if e != nil {
		t.Error(e)
		return
	}

	e = Put(`gs://testBucket/36_3p53hr4z.log`, r)
	if e != nil {
		t.Error(e)
	}
}

type CC struct {
}

func (c CC) Close() error {
	return errors.New(`test error`)
}
func TestCC(t *testing.T) {
	checkClose(CC{})
}
