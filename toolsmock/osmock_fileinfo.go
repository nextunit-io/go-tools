package toolsmock

import (
	"fmt"
	"os"
	"time"

	gomock "github.com/nextunit-io/go-mock"
)

type fileInfoMockStruct struct {
	Name *gomock.ToolMock[
		interface{},
		string,
	]
	Size *gomock.ToolMock[
		interface{},
		int64,
	]
	Mode *gomock.ToolMock[
		interface{},
		os.FileMode,
	]
	ModTime *gomock.ToolMock[
		interface{},
		time.Time,
	]
	IsDir *gomock.ToolMock[
		interface{},
		bool,
	]
	Sys *gomock.ToolMock[
		interface{},
		any,
	]
}

type FileInfoMock struct {
	Mock fileInfoMockStruct
}

func GetFileInfoMock() *FileInfoMock {
	return &FileInfoMock{
		Mock: fileInfoMockStruct{
			Name: gomock.GetMock[
				interface{},
				string,
			](fmt.Errorf("Name general error")),
			Size: gomock.GetMock[
				interface{},
				int64,
			](fmt.Errorf("Size general error")),
			Mode: gomock.GetMock[
				interface{},
				os.FileMode,
			](fmt.Errorf("Mode general error")),
			ModTime: gomock.GetMock[
				interface{},
				time.Time,
			](fmt.Errorf("ModTime general error")),
			IsDir: gomock.GetMock[
				interface{},
				bool,
			](fmt.Errorf("IsDir general error")),
			Sys: gomock.GetMock[
				interface{},
				any,
			](fmt.Errorf("Sys general error")),
		},
	}
}

func (fileInfoMock *FileInfoMock) Name() string {
	fileInfoMock.Mock.Name.AddInput(nil)

	result, err := fileInfoMock.Mock.Name.GetNextResult()
	if err != nil {
		panic(err.Error())
	}

	return *result
}

func (fileInfoMock *FileInfoMock) Size() int64 {
	fileInfoMock.Mock.Size.AddInput(nil)

	result, err := fileInfoMock.Mock.Size.GetNextResult()
	if err != nil {
		panic(err.Error())
	}

	return *result
}

func (fileInfoMock *FileInfoMock) Mode() os.FileMode {
	fileInfoMock.Mock.Mode.AddInput(nil)

	result, err := fileInfoMock.Mock.Mode.GetNextResult()
	if err != nil {
		panic(err.Error())
	}

	return *result
}

func (fileInfoMock *FileInfoMock) ModTime() time.Time {
	fileInfoMock.Mock.ModTime.AddInput(nil)

	result, err := fileInfoMock.Mock.ModTime.GetNextResult()
	if err != nil {
		panic(err.Error())
	}

	return *result
}

func (fileInfoMock *FileInfoMock) IsDir() bool {
	fileInfoMock.Mock.IsDir.AddInput(nil)

	result, err := fileInfoMock.Mock.IsDir.GetNextResult()
	if err != nil {
		panic(err.Error())
	}

	return *result
}

func (fileInfoMock *FileInfoMock) Sys() any {
	fileInfoMock.Mock.Sys.AddInput(nil)

	result, err := fileInfoMock.Mock.Sys.GetNextResult()
	if err != nil {
		panic(err.Error())
	}

	return *result
}
