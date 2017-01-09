package fisherking

import (
	"context"
	"io/ioutil"
	"log"
	"testing"
)

func TestReadS3(t *testing.T) {
	cxt := context.Background()
	r, e := s3{cxt}.Get(`s3://my/53hr4z.log`)
	if e != nil {
		t.Error(e)
		return
	}

	d, e := ioutil.ReadAll(r)
	if e != nil {
		t.Log(e)
		t.Fail()
	}
	log.Printf("%s", d)
}
