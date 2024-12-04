package toolsmock

import (
	"fmt"
	"os/exec"

	gomock "github.com/nextunit-io/go-mock"
)

type execMockStruct struct {
	Command *gomock.ToolMock[
		struct {
			Name string
			Arg  []string
		},
		*exec.Cmd,
	]
}

type ExecMock struct {
	Mock execMockStruct
}

func GetExecMock() *ExecMock {
	return &ExecMock{
		Mock: execMockStruct{
			Command: gomock.GetMock[
				struct {
					Name string
					Arg  []string
				},
				*exec.Cmd,
			](fmt.Errorf("COMMAND general error")),
		},
	}
}

func (e *ExecMock) Command(name string, arg ...string) *exec.Cmd {
	e.Mock.Command.AddInput(struct {
		Name string
		Arg  []string
	}{
		Name: name,
		Arg:  arg,
	})

	result, err := e.Mock.Command.GetNextResult()
	if err != nil {
		panic(err.Error())
	}

	return *result
}
