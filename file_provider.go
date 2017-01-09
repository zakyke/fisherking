package fisherking

import (
	"context"
	"errors"
	"io"
	"log"
	"path"

	"os"
)

type fs struct {
	context.Context
}

func (f fs) Get(path string) (io.Reader, error) {
	path = path[len(fsPrefix):]
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (f fs) Put(destination string, data io.Reader) error {
	// Create directory if needed.
	destination = destination[len(fsPrefix):]
	basepath := path.Dir(destination)
	filename := path.Base(destination)
	log.Println(basepath, ` `, filename)
	if os.MkdirAll(basepath, 0700) != nil {
		return errors.New(`unable to create directory ` + basepath)

	}

	fileOut, err := os.Create(path.Join(basepath, filename))
	if err != nil {
		return errors.New(`nable to create ` + path.Join(basepath, filename) + err.Error())
	}
	defer fileOut.Close()
	return nil
}
