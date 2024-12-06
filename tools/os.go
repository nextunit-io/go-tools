package tools

import (
	"os"
)

type OsInterface interface {
	MkdirAll(path string, perm os.FileMode) error
	MkdirTemp(dir, pattern string) (string, error)
	Mkdir(name string, perm os.FileMode) error
	ReadFile(name string) ([]byte, error)
	Remove(name string) error
	RemoveAll(path string) error
}

type defaultOsClient struct {
	OsInterface
}

var osInstane OsInterface

func GetOsInstance() OsInterface {
	if osInstane == nil {
		osInstane = &defaultOsClient{}
	}

	return osInstane
}

func SetOsInstance(client OsInterface) {
	osInstane = client
}

func (defaultOsClient) MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

func (defaultOsClient) MkdirTemp(dir, pattern string) (string, error) {
	return os.MkdirTemp(dir, pattern)
}
func (defaultOsClient) Mkdir(name string, perm os.FileMode) error {
	return os.Mkdir(name, perm)
}
func (defaultOsClient) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}
func (defaultOsClient) Remove(name string) error {
	return os.Remove(name)
}
func (defaultOsClient) RemoveAll(path string) error {
	return os.RemoveAll(path)
}
