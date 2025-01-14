package toolsmock

import (
	"fmt"
	"os"

	gomock "github.com/nextunit-io/go-mock"
)

type dirEntryMockStruct struct {
	Info *gomock.ToolMock[
		interface{},
		os.FileInfo,
	]
	IsDir *gomock.ToolMock[
		interface{},
		bool,
	]
	Name *gomock.ToolMock[
		interface{},
		string,
	]
	Type *gomock.ToolMock[
		interface{},
		os.FileMode,
	]
}

type DirEntryMock struct {
	Mock dirEntryMockStruct
}

func GetDirEntryMock() *DirEntryMock {
	return &DirEntryMock{
		Mock: dirEntryMockStruct{
			Info: gomock.GetMock[
				interface{},
				os.FileInfo,
			](fmt.Errorf("Info general error")),
			IsDir: gomock.GetMock[
				interface{},
				bool,
			](fmt.Errorf("IsDir general error")),
			Name: gomock.GetMock[
				interface{},
				string,
			](fmt.Errorf("Name general error")),
			Type: gomock.GetMock[
				interface{},
				os.FileMode,
			](fmt.Errorf("Type general error")),
		},
	}
}

func (dirEntryMock *DirEntryMock) Info() (os.FileInfo, error) {
	dirEntryMock.Mock.Info.AddInput(struct{}{})

	result, err := dirEntryMock.Mock.Info.GetNextResult()
	if err != nil {
		return nil, err
	}

	return *result, nil
}

func (dirEntryMock *DirEntryMock) IsDir() bool {
	dirEntryMock.Mock.IsDir.AddInput(struct{}{})

	result, err := dirEntryMock.Mock.IsDir.GetNextResult()
	if err != nil {
		panic(err)
	}

	return *result
}

func (dirEntryMock *DirEntryMock) Name() string {
	dirEntryMock.Mock.Name.AddInput(struct{}{})

	result, _ := dirEntryMock.Mock.Name.GetNextResult()
	return *result
}

func (dirEntryMock *DirEntryMock) Type() os.FileMode {
	dirEntryMock.Mock.Type.AddInput(struct{}{})

	result, err := dirEntryMock.Mock.Type.GetNextResult()
	if err != nil {
		panic(err)
	}

	return *result
}
