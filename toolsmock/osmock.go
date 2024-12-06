package toolsmock

import (
	"fmt"
	"os"

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
