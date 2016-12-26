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

func TestRead(t *testing.T) {
	r, e := Get(`file:///tmp/list.txt`)
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
