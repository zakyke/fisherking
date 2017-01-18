package fisherking

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	largeContentStringLength = 8 //len(10000000)  = 10MB in  bytes
)

type htp struct {
	context.Context
}

func (h htp) Get(path string) (io.Reader, error) {
	hresp, err := http.Head(path)
	if err != nil {
		return nil, err
	}

	if l := len(hresp.Header.Get(`Content-Length`)); l >= largeContentStringLength {
		return h.pipeGet(path)
	}
	return h.get(path)
}

func (h htp) pipeGet(path string) (io.Reader, error) {
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	if h := h.headers(); h != nil {
		req.Header = h
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	r, w := io.Pipe()

	go func() {
		defer checkClose(w)
		defer checkClose(res.Body)
		_, e := io.Copy(w, res.Body)
		if e != nil {
			log.Println(e)
		}
	}()

	return r, err
}

func (h htp) get(path string) (io.Reader, error) {
	//TODO Check for content length and decide wether to use a pipe or in memory buffer.

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	if h := h.headers(); h != nil {
		req.Header = h
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer checkClose(res.Body)
	d, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(d), err
}

func (h htp) headers() http.Header {
	if val := h.Context.Value(`http.headers`); val != nil {
		if h, ok := val.(http.Header); ok {
			return h
		}
	}
	return nil
}

func (h htp) Put(destination string, data io.Reader) error {
	return nil
}
