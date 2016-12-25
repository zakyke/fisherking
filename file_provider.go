package fisherking

import (
	"context"
	"io"

	"os"
)

type fs struct {
	context.Context
}

func (f fs) GetWithContext(cxt context.Context, path string) FileGetter {
	return fs{cxt}.Get
}

func (f fs) WithReadAll() FileGetter {
	return fs{}.Get
}
func (f fs) Get(path string) (io.Reader, error) {
	path = path[6:]
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// type AllReader func(path string) []byte

// func (getter AllReader) ReadAll(fr FileGetter) []byte {
// 	r, e := fr(getter.)
// 	if e != nil {
// 		log.Print(e)
// 		return []byte{}
// 	}
// 	d, e := ioutil.ReadAll(r)
// 	if e != nil {
// 		return []byte{}
// 	}
// 	return d
// }

func (f fs) PutWithContext(cxt context.Context, path string) FilePutter {
	return fs{cxt}.Put
}
func (f fs) Put(destination string) (io.Writer, error) {
	file, err := os.Open(destination)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// func getFile(path string) (io.Reader, error) {
// 	return nil, nil
// }

// func (gf FileGetter) WithContext(context context.Context) FileGetter {
// 	return getFile
// }

// func parser(path string) (string, string) {
// 	// if runtime.GOOS == `windows` {
// 	// 	return "", parseWinFS(path)
// 	// }
// 	// return "", parseLnxFS(path)
// 	//return "", parseLnxFS(path)
// }
