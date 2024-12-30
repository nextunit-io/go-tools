package tools

import (
	"os"

	"github.com/nextunit-io/go-tools/interfaces"
)

type OsInterface interface {
	Create(name string) (interfaces.OsFileInterface, error)
	MkdirAll(path string, perm os.FileMode) error
	MkdirTemp(dir, pattern string) (string, error)
	Mkdir(name string, perm os.FileMode) error
	Open(name string) (interfaces.OsFileInterface, error)
	OpenFile(name string, flag int, perm os.FileMode) (interfaces.OsFileInterface, error)
	ReadFile(name string) ([]byte, error)
	Remove(name string) error
	RemoveAll(path string) error
	Stat(name string) (os.FileInfo, error)
	TempDir() string
	UserHomeDir() (string, error)
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

func (defaultOsClient) Create(name string) (interfaces.OsFileInterface, error) {
	return os.Create(name)
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
func (defaultOsClient) Open(name string) (interfaces.OsFileInterface, error) {
	return os.Open(name)
}
func (defaultOsClient) OpenFile(name string, flag int, perm os.FileMode) (interfaces.OsFileInterface, error) {
	return os.OpenFile(name, flag, perm)
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
func (defaultOsClient) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}
func (defaultOsClient) TempDir() string {
	return os.TempDir()
}
func (defaultOsClient) UserHomeDir() (string, error) {
	return os.UserHomeDir()
}
