package toolsmock

import (
	"fmt"
	"os"
	"time"

	gomock "github.com/nextunit-io/go-mock"
)

type osMockStruct struct {
	MkdirAll *gomock.ToolMock[
		struct {
			Path string
			Perm os.FileMode
		},
		bool,
	]
	MkdirTemp *gomock.ToolMock[
		struct {
			Dir     string
			Pattern string
		},
		string,
	]
	Mkdir *gomock.ToolMock[
		struct {
			Name string
			Perm os.FileMode
		},
		bool,
	]
	ReadFile *gomock.ToolMock[
		struct {
			Name string
		},
		[]byte,
	]
	Remove *gomock.ToolMock[
		struct {
			Name string
		},
		bool,
	]
	RemoveAll *gomock.ToolMock[
		struct {
			Path string
		},
		bool,
	]
	Stat *gomock.ToolMock[
		struct {
			Name string
		},
		os.FileInfo,
	]
	TempDir *gomock.ToolMock[
		interface{},
		string,
	]
	UserHomeDir *gomock.ToolMock[
		interface{},
		string,
	]
}

type OsMock struct {
	Mock osMockStruct
}

func GetOsMock() *OsMock {
	return &OsMock{
		Mock: osMockStruct{
			MkdirAll: gomock.GetMock[
				struct {
					Path string
					Perm os.FileMode
				},
				bool,
			](fmt.Errorf("MkdirAll general error")),
			MkdirTemp: gomock.GetMock[
				struct {
					Dir     string
					Pattern string
				},
				string,
			](fmt.Errorf("MkdirTemp general error")),
			Mkdir: gomock.GetMock[
				struct {
					Name string
					Perm os.FileMode
				},
				bool,
			](fmt.Errorf("Mkdir general error")),
			ReadFile: gomock.GetMock[
				struct {
					Name string
				},
				[]byte,
			](fmt.Errorf("ReadFile general error")),
			Remove: gomock.GetMock[
				struct {
					Name string
				},
				bool,
			](fmt.Errorf("Remove general error")),
			RemoveAll: gomock.GetMock[
				struct {
					Path string
				},
				bool,
			](fmt.Errorf("RemoveAll general error")),
			Stat: gomock.GetMock[
				struct {
					Name string
				},
				os.FileInfo,
			](fmt.Errorf("Stat general error")),
			TempDir: gomock.GetMock[
				interface{},
				string,
			](fmt.Errorf("TempDir general error")),
			UserHomeDir: gomock.GetMock[
				interface{},
				string,
			](fmt.Errorf("UserHomeDir general error")),
		},
	}
}

func (osMock *OsMock) MkdirAll(path string, perm os.FileMode) error {
	osMock.Mock.MkdirAll.AddInput(struct {
		Path string
		Perm os.FileMode
	}{
		Path: path,
		Perm: perm,
	})

	result, err := osMock.Mock.MkdirAll.GetNextResult()
	if err != nil {
		return err
	}
	if *result {
		return fmt.Errorf("Error requested")
	}

	return nil
}

func (osMock *OsMock) MkdirTemp(dir, pattern string) (string, error) {
	osMock.Mock.MkdirTemp.AddInput(struct {
		Dir     string
		Pattern string
	}{
		Dir:     dir,
		Pattern: pattern,
	})

	result, err := osMock.Mock.MkdirTemp.GetNextResult()
	if err != nil {
		return "", err
	}

	return *result, nil
}

func (osMock *OsMock) Mkdir(name string, perm os.FileMode) error {
	osMock.Mock.Mkdir.AddInput(struct {
		Name string
		Perm os.FileMode
	}{
		Name: name,
		Perm: perm,
	})

	result, err := osMock.Mock.Mkdir.GetNextResult()
	if err != nil {
		return err
	}
	if *result {
		return fmt.Errorf("Error requested")
	}

	return nil
}

func (osMock *OsMock) ReadFile(name string) ([]byte, error) {
	osMock.Mock.ReadFile.AddInput(struct {
		Name string
	}{
		Name: name,
	})

	result, err := osMock.Mock.ReadFile.GetNextResult()
	if err != nil {
		return nil, err
	}

	return *result, nil
}

func (osMock *OsMock) Remove(name string) error {
	osMock.Mock.Remove.AddInput(struct {
		Name string
	}{
		Name: name,
	})

	result, err := osMock.Mock.Remove.GetNextResult()
	if err != nil {
		return err
	}
	if *result {
		return fmt.Errorf("Error requested")
	}

	return nil
}

func (osMock *OsMock) RemoveAll(path string) error {
	osMock.Mock.RemoveAll.AddInput(struct {
		Path string
	}{
		Path: path,
	})

	result, err := osMock.Mock.RemoveAll.GetNextResult()
	if err != nil {
		return err
	}
	if *result {
		return fmt.Errorf("Error requested")
	}

	return nil
}

func (osMock *OsMock) Stat(name string) (os.FileInfo, error) {
	osMock.Mock.Stat.AddInput(struct {
		Name string
	}{
		Name: name,
	})

	result, err := osMock.Mock.Stat.GetNextResult()
	if err != nil {
		return nil, err
	}

	return *result, nil
}

func (osMock *OsMock) TempDir() string {
	osMock.Mock.TempDir.AddInput(nil)

	result, err := osMock.Mock.TempDir.GetNextResult()
	if err != nil {
		panic(err.Error())
	}

	return *result
}

func (osMock *OsMock) UserHomeDir() string {
	osMock.Mock.UserHomeDir.AddInput(nil)

	result, err := osMock.Mock.UserHomeDir.GetNextResult()
	if err != nil {
		panic(err.Error())
	}

	return *result
}

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
