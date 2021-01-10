package filesys

import (
	"fmt"
	"io/ioutil"
	"os"
)

// FileStat for monkey patch
var FileStat = os.Stat
var ReadFile = ioutil.ReadFile

type IFileSys interface {
	Exists(string) bool
	ReadAll(string) ([]byte, error)
}

// File wrapper
type File struct{}

// New creates a file wrapper
func New() *File {
	return &File{}
}

// Exists returns true if file exists. Otherwise returns false.
func (f *File) Exists(filename string) bool {
	if _, err := FileStat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

// ReadAll returns file content in []byte with corresponding error
func (f *File) ReadAll(filename string) ([]byte, error) {
	if !f.Exists(filename) {
		return nil, fmt.Errorf("File %s does not exist", filename)
	}

	bytes, err := ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
