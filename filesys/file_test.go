package filesys_test

import (
	"dataman/filesys"
	"errors"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileExistsShouldReturnTrue(t *testing.T) {
	defer (func() {
		filesys.FileStat = os.Stat
	})()

	filesys.FileStat = func(filename string) (os.FileInfo, error) {
		return nil, nil
	}

	f := filesys.New()
	result := f.Exists("an_existing_file")

	assert.True(t, result)
}

func TestFileDoesNotExistShouldReturnFalse(t *testing.T) {
	defer (func() {
		filesys.FileStat = os.Stat
	})()

	filesys.FileStat = func(filename string) (os.FileInfo, error) {
		return nil, os.ErrNotExist
	}

	f := filesys.New()
	result := f.Exists("a_non_existing_file")

	assert.False(t, result)
}

func TestReadNonExistingFileShouldReturnError(t *testing.T) {
	defer (func() {
		filesys.FileStat = os.Stat
	})()

	filesys.FileStat = func(filename string) (os.FileInfo, error) {
		return nil, os.ErrNotExist
	}

	f := filesys.New()
	data, err := f.ReadAll("a_non_existing_file")

	assert.Nil(t, data)
	assert.NotNil(t, err)
}

func TestReadExistingFileShouldReturnData(t *testing.T) {
	defer (func() {
		filesys.FileStat = os.Stat
		filesys.ReadFile = ioutil.ReadFile
	})()

	expectedFileContent := []byte{1, 2, 3, 4}

	filesys.FileStat = func(filename string) (os.FileInfo, error) {
		return nil, nil
	}

	filesys.ReadFile = func(filename string) ([]byte, error) {
		return expectedFileContent, nil
	}

	f := filesys.New()
	data, err := f.ReadAll("an_existing_file")

	assert.Equal(t, expectedFileContent, data)
	assert.Nil(t, err)
}

func TestReadInaccessibleExistingFileShouldReturnError(t *testing.T) {
	defer (func() {
		filesys.FileStat = os.Stat
		filesys.ReadFile = ioutil.ReadFile
	})()

	filesys.FileStat = func(filename string) (os.FileInfo, error) {
		return nil, nil
	}

	filesys.ReadFile = func(filename string) ([]byte, error) {
		return nil, errors.New("you cannot read this file for some reason")
	}

	f := filesys.New()
	data, err := f.ReadAll("an_existing_file")

	assert.Nil(t, data)
	assert.NotNil(t, err)
}
