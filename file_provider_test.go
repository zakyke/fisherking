package fisherking

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestReadFS(t *testing.T) {
	r, e := fs{}.Get(`file:///tmp/list.txt`)
	if e != nil {
		t.Log(e)
		t.Fail()
	}
	d, e := ioutil.ReadAll(r)
	if e != nil {
		t.Log(e)
		t.Fail()
	}
	log.Printf("%s", d)
}

func TestReadLocalPutGCS(t *testing.T) {
	object := `2017/01/01/10/56f91cd4d3e3660002000033/adsmanager/http-collector-group-us-multi-preemp-8qdf-adsmanager-1483228909992095856.gz`
	//get file
	r, e := Get(`file:///home/zaky/streamrail/go/src/github.com/streamrail/messaging/data/` + object)
	if e != nil {
		t.Log(e)
		t.Fail()
	}
	//put file
	e = Put(`gs://testtc/`+object, r)
	if e != nil {
		t.Log(e)
		t.Fail()
	}
}

func TestReadGCSPutLocal(t *testing.T) {
	object := `2017/01/01/10/56f91cd4d3e3660002000033/adsmanager/http-collector-group-eu-3v9n-adsmanager-1483229134317833261.gz`
	//get file
	r, e := Get(`gs://testtc/` + object)
	if e != nil {
		t.Log(e)
		t.Fail()
	}
	//put file
	e = Put(`file:///tmp/github.com/streamrail/messaging/data/`+object+`-fromGS`, r)
	if e != nil {
		t.Log(e)
		t.Fail()
	}
}
