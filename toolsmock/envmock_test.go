package toolsmock_test

import (
	"fmt"
	"testing"

	"github.com/nextunit-io/go-tools/toolsmock"
	"github.com/stretchr/testify/assert"
)

func TestEnvMockGetenv(t *testing.T) {
	t.Helper()
	envMock := toolsmock.GetEnvMock()

	t.Run("Testing Getenv", func(t *testing.T) {
		output := "test-string-output"

		envMock.Mock.Getenv.AddReturnValue(&output)
		envMock.Mock.Getenv.AddReturnValue(&output)
		envMock.Mock.Getenv.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			o := envMock.Getenv(fmt.Sprintf("test-input-%d", i))
			assert.Equal(t, output, o)
		}

		assert.Equal(t, 3, envMock.Mock.Getenv.HasBeenCalled())
		for i := 0; i < 3; i++ {
			input := envMock.Mock.Getenv.GetInput(i)
			assert.Equal(t, fmt.Sprintf("test-input-%d", i), input.Key)
		}
	})
}
