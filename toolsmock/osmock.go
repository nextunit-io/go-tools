package toolsmock

import (
	"fmt"
	"os"

	gomock "github.com/nextunit-io/go-mock"
	"github.com/nextunit-io/go-tools/interfaces"
)

type osMockStruct struct {
	Create *gomock.ToolMock[
		struct {
			Name string
		},
		interfaces.OsFileInterface,
	]
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
	OpenFile *gomock.ToolMock[
		struct {
			Name string
			Flag int
			Perm os.FileMode
		},
		interfaces.OsFileInterface,
	]
	ReadDir *gomock.ToolMock[
		struct {
			Name string
		},
		[]os.DirEntry,
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
	Open *gomock.ToolMock[
		struct {
			Name string
		},
		interfaces.OsFileInterface,
	]
}

type OsMock struct {
	Mock osMockStruct
}

func GetOsMock() *OsMock {
	return &OsMock{
		Mock: osMockStruct{
			Create: gomock.GetMock[
				struct {
					Name string
				},
				interfaces.OsFileInterface,
			](fmt.Errorf("Create general error")),
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
			Open: gomock.GetMock[
				struct {
					Name string
				},
				interfaces.OsFileInterface,
			](fmt.Errorf("Open general error")),
			OpenFile: gomock.GetMock[
				struct {
					Name string
					Flag int
					Perm os.FileMode
				},
				interfaces.OsFileInterface,
			](fmt.Errorf("OpenFile general error")),
			ReadDir: gomock.GetMock[
				struct {
					Name string
				},
				[]os.DirEntry,
			](fmt.Errorf("ReadDir general error")),
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

func (osMock *OsMock) Create(name string) (interfaces.OsFileInterface, error) {
	osMock.Mock.Create.AddInput(struct {
		Name string
	}{
		name,
	})

	result, err := osMock.Mock.Create.GetNextResult()
	if err != nil {
		return nil, err
	}

	return *result, nil
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

func (osMock *OsMock) OpenFile(name string, flag int, perm os.FileMode) (interfaces.OsFileInterface, error) {
	osMock.Mock.OpenFile.AddInput(struct {
		Name string
		Flag int
		Perm os.FileMode
	}{
		Name: name,
		Flag: flag,
		Perm: perm,
	})

	result, err := osMock.Mock.OpenFile.GetNextResult()
	if err != nil {
		return nil, err
	}

	return *result, nil
}

func (osMock *OsMock) ReadDir(name string) ([]os.DirEntry, error) {
	osMock.Mock.ReadDir.AddInput(struct {
		Name string
	}{
		Name: name,
	})

	result, err := osMock.Mock.ReadDir.GetNextResult()
	if err != nil {
		return nil, err
	}

	return *result, nil
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

func (osMock *OsMock) UserHomeDir() (string, error) {
	osMock.Mock.UserHomeDir.AddInput(nil)

	result, err := osMock.Mock.UserHomeDir.GetNextResult()
	if err != nil {
		return "", err
	}

	return *result, nil
}

func (osMock *OsMock) Open(name string) (interfaces.OsFileInterface, error) {
	osMock.Mock.Open.AddInput(struct {
		Name string
	}{
		Name: name,
	})

	result, err := osMock.Mock.Open.GetNextResult()
	if err != nil {
		return nil, err
	}

	return *result, nil
}
