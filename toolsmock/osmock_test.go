package toolsmock_test

import (
	"fmt"
	"io/fs"
	"testing"

	"github.com/nextunit-io/go-tools/toolsmock"
	"github.com/stretchr/testify/assert"
)

func TestOsMockMkdirAll(t *testing.T) {
	t.Helper()
	osMock := toolsmock.GetOsMock()

	t.Run("Testing MkdirAll", func(t *testing.T) {
		mkdirErrorReturn := false
		perm := fs.ModeAppend.Perm()

		osMock.Mock.MkdirAll.AddReturnValue(&mkdirErrorReturn)
		osMock.Mock.MkdirAll.AddReturnValue(&mkdirErrorReturn)
		osMock.Mock.MkdirAll.AddReturnValue(&mkdirErrorReturn)

		for i := 0; i < 3; i++ {
			err := osMock.MkdirAll(fmt.Sprintf("test-name-%d", i), perm)
			assert.Nil(t, err)
		}

		err := osMock.MkdirAll("test-input-error", perm)
		assert.Equal(t, fmt.Errorf("MkdirAll general error"), err)

		assert.Equal(t, 4, osMock.Mock.MkdirAll.HasBeenCalled())

		for i := 0; i < 3; i++ {
			input := osMock.Mock.MkdirAll.GetInput(i)
			assert.Equal(t, fmt.Sprintf("test-name-%d", i), input.Path)
			assert.Equal(t, perm, input.Perm)
		}

		input := osMock.Mock.MkdirAll.GetInput(3)
		assert.Equal(t, "test-input-error", input.Path)
		assert.Equal(t, perm, input.Perm)
	})
}

func TestOsMockMkdirTemp(t *testing.T) {
	t.Helper()
	osMock := toolsmock.GetOsMock()

	t.Run("Testing MkdirTemp", func(t *testing.T) {
		mkdirTempReturn := "test-path"

		osMock.Mock.MkdirTemp.AddReturnValue(&mkdirTempReturn)
		osMock.Mock.MkdirTemp.AddReturnValue(&mkdirTempReturn)
		osMock.Mock.MkdirTemp.AddReturnValue(&mkdirTempReturn)

		for i := 0; i < 3; i++ {
			path, err := osMock.MkdirTemp(fmt.Sprintf("test-name-%d", i), fmt.Sprintf("test-pattern-%d", i))
			assert.Equal(t, mkdirTempReturn, path)
			assert.Nil(t, err)
		}

		path, err := osMock.MkdirTemp("test-input-error", "test-pattern-error")
		assert.Equal(t, "", path)
		assert.Equal(t, fmt.Errorf("MkdirTemp general error"), err)

		assert.Equal(t, 4, osMock.Mock.MkdirTemp.HasBeenCalled())
		for i := 0; i < 3; i++ {
			input := osMock.Mock.MkdirTemp.GetInput(i)
			assert.Equal(t, fmt.Sprintf("test-name-%d", i), input.Dir)
			assert.Equal(t, fmt.Sprintf("test-pattern-%d", i), input.Pattern)
		}

		input := osMock.Mock.MkdirTemp.GetInput(3)
		assert.Equal(t, "test-input-error", input.Dir)
		assert.Equal(t, "test-pattern-error", input.Pattern)
	})
}

func TestOsMockMkdir(t *testing.T) {
	t.Helper()
	osMock := toolsmock.GetOsMock()

	t.Run("Testing Mkdir", func(t *testing.T) {
		mkdirErrorReturn := false
		perm := fs.ModeAppend.Perm()

		osMock.Mock.Mkdir.AddReturnValue(&mkdirErrorReturn)
		osMock.Mock.Mkdir.AddReturnValue(&mkdirErrorReturn)
		osMock.Mock.Mkdir.AddReturnValue(&mkdirErrorReturn)

		for i := 0; i < 3; i++ {
			err := osMock.Mkdir(fmt.Sprintf("test-name-%d", i), perm)
			assert.Nil(t, err)
		}

		err := osMock.Mkdir("test-input-error", perm)
		assert.Equal(t, fmt.Errorf("Mkdir general error"), err)

		assert.Equal(t, 4, osMock.Mock.Mkdir.HasBeenCalled())

		for i := 0; i < 3; i++ {
			input := osMock.Mock.Mkdir.GetInput(i)
			assert.Equal(t, fmt.Sprintf("test-name-%d", i), input.Name)
			assert.Equal(t, perm, input.Perm)
		}

		input := osMock.Mock.Mkdir.GetInput(3)
		assert.Equal(t, "test-input-error", input.Name)
		assert.Equal(t, perm, input.Perm)
	})
}

func TestOsMockReadFile(t *testing.T) {
	t.Helper()
	osMock := toolsmock.GetOsMock()

	t.Run("Testing ReadFile", func(t *testing.T) {
		readFileReturn := []byte("test-content")

		osMock.Mock.ReadFile.AddReturnValue(&readFileReturn)
		osMock.Mock.ReadFile.AddReturnValue(&readFileReturn)
		osMock.Mock.ReadFile.AddReturnValue(&readFileReturn)

		for i := 0; i < 3; i++ {
			content, err := osMock.ReadFile(fmt.Sprintf("test-name-%d", i))
			assert.Equal(t, readFileReturn, content)
			assert.Nil(t, err)
		}

		content, err := osMock.ReadFile("test-input-error")
		assert.Nil(t, content)
		assert.Equal(t, fmt.Errorf("ReadFile general error"), err)

		assert.Equal(t, 4, osMock.Mock.ReadFile.HasBeenCalled())
		for i := 0; i < 3; i++ {
			input := osMock.Mock.ReadFile.GetInput(i)
			assert.Equal(t, fmt.Sprintf("test-name-%d", i), input.Name)
		}

		input := osMock.Mock.ReadFile.GetInput(3)
		assert.Equal(t, "test-input-error", input.Name)
	})
}

func TestOsMockRemove(t *testing.T) {
	t.Helper()
	osMock := toolsmock.GetOsMock()

	t.Run("Testing Remove", func(t *testing.T) {
		RemoveErrorReturn := false

		osMock.Mock.Remove.AddReturnValue(&RemoveErrorReturn)
		osMock.Mock.Remove.AddReturnValue(&RemoveErrorReturn)
		osMock.Mock.Remove.AddReturnValue(&RemoveErrorReturn)

		for i := 0; i < 3; i++ {
			err := osMock.Remove(fmt.Sprintf("test-name-%d", i))
			assert.Nil(t, err)
		}

		err := osMock.Remove("test-input-error")
		assert.Equal(t, fmt.Errorf("Remove general error"), err)

		assert.Equal(t, 4, osMock.Mock.Remove.HasBeenCalled())

		for i := 0; i < 3; i++ {
			input := osMock.Mock.Remove.GetInput(i)
			assert.Equal(t, fmt.Sprintf("test-name-%d", i), input.Name)
		}

		input := osMock.Mock.Remove.GetInput(3)
		assert.Equal(t, "test-input-error", input.Name)
	})
}

func TestOsMockRemoveAll(t *testing.T) {
	t.Helper()
	osMock := toolsmock.GetOsMock()

	t.Run("Testing RemoveAll", func(t *testing.T) {
		RemoveAllErrorReturn := false

		osMock.Mock.RemoveAll.AddReturnValue(&RemoveAllErrorReturn)
		osMock.Mock.RemoveAll.AddReturnValue(&RemoveAllErrorReturn)
		osMock.Mock.RemoveAll.AddReturnValue(&RemoveAllErrorReturn)

		for i := 0; i < 3; i++ {
			err := osMock.RemoveAll(fmt.Sprintf("test-name-%d", i))
			assert.Nil(t, err)
		}

		err := osMock.RemoveAll("test-input-error")
		assert.Equal(t, fmt.Errorf("RemoveAll general error"), err)

		assert.Equal(t, 4, osMock.Mock.RemoveAll.HasBeenCalled())

		for i := 0; i < 3; i++ {
			input := osMock.Mock.RemoveAll.GetInput(i)
			assert.Equal(t, fmt.Sprintf("test-name-%d", i), input.Path)
		}

		input := osMock.Mock.RemoveAll.GetInput(3)
		assert.Equal(t, "test-input-error", input.Path)
	})
}
