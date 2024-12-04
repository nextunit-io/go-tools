package tools

import (
	"os/exec"
)

type ExecInterface interface {
	Command(name string, arg ...string) *exec.Cmd
}

type defaultExecClient struct {
	ExecInterface
}

var execInstane ExecInterface

func GetExecInstance() ExecInterface {
	if execInstane == nil {
		execInstane = &defaultExecClient{}
	}

	return execInstane
}

func SetExecInstance(client ExecInterface) {
	execInstane = client
}

func (defaultExecClient) Command(name string, arg ...string) *exec.Cmd {
	return exec.Command(name, arg...)
}
