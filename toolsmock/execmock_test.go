package toolsmock_test

import (
	"fmt"
	"os/exec"
	"testing"

	"github.com/nextunit-io/go-tools/toolsmock"
	"github.com/stretchr/testify/assert"
)

func TestExecMockCommand(t *testing.T) {
	t.Helper()
	execMock := toolsmock.GetExecMock()

	t.Run("Testing Command", func(t *testing.T) {
		cmdReturn := &exec.Cmd{
			Path: "test-path",
		}

		execMock.Mock.Command.AddReturnValue(&cmdReturn)
		execMock.Mock.Command.AddReturnValue(&cmdReturn)
		execMock.Mock.Command.AddReturnValue(&cmdReturn)

		for i := 0; i < 3; i++ {
			cmd := execMock.Command(fmt.Sprintf("test-name-%d", i), "arg1", "arg2")
			assert.Equal(t, cmdReturn, cmd)
		}

		assert.Equal(t, 3, execMock.Mock.Command.HasBeenCalled())
		for i := 0; i < 3; i++ {
			input := execMock.Mock.Command.GetInput(i)
			assert.Equal(t, fmt.Sprintf("test-name-%d", i), input.Name)
			assert.Equal(t, []string{"arg1", "arg2"}, input.Arg)
		}

		execMock.Mock.Command.Reset()
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}

			assert.Equal(t, 1, execMock.Mock.Command.HasBeenCalled())
			input := execMock.Mock.Command.GetInput(0)
			assert.Equal(t, "test-name-error", input.Name)
			assert.Equal(t, []string{"arg1", "arg2"}, input.Arg)
		}()

		_ = execMock.Command("test-name-error", "arg1", "arg2")
	})
}
