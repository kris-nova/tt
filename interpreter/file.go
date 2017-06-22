package interpreter

import (
	"io/ioutil"
	"path/filepath"
)

type File struct {
	AbsolutePath string
	Bytes        []byte
}

func LoadFile(path string) (*File, error) {
	f := &File{}
	abspath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	f.AbsolutePath = abspath
	bytes, err := ioutil.ReadFile(abspath)
	if err != nil {
		return nil, err
	}
	f.Bytes = bytes
	return f, nil
}

func (f *File) String() string {
	return string(f.Bytes)
}
