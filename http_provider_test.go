package fisherking

import (
	"context"
	"log"
	"net/http"
)

func Examplehtp_Get() {
	var hd http.Header
	hd.Add(`Authorization`, `usr:passwd`)
	var hk interface{} = `http.headers`
	cxt := context.WithValue(context.Background(), hk, hd)
	r, e := GetWithContext(cxt, `http://releases.ubuntu.com/16.04.1/ubuntu-16.04.1-desktop-amd64.iso?_ga=1.185584303.702866013.1483453284`)
	if e != nil {
		log.Println(e)
		return
	}
	log.Println(r)
}
