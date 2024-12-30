package toolsmock_test

import (
	"fmt"
	"io/fs"
	"os"
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

func TestOsMockOpen(t *testing.T) {
	t.Helper()
	osMock := toolsmock.GetOsMock()

	t.Run("Testing Open", func(t *testing.T) {
		var openReturn *os.File = &os.File{}

		osMock.Mock.Open.AddReturnValue(&openReturn)
		osMock.Mock.Open.AddReturnValue(&openReturn)
		osMock.Mock.Open.AddReturnValue(&openReturn)

		for i := 0; i < 3; i++ {
			file, err := osMock.Open(fmt.Sprintf("test-name-%d", i))
			assert.Equal(t, openReturn, file)
			assert.Nil(t, err)
		}

		file, err := osMock.Open("test-input-error")
		assert.Nil(t, file)
		assert.Equal(t, fmt.Errorf("Open general error"), err)

		assert.Equal(t, 4, osMock.Mock.Open.HasBeenCalled())

		for i := 0; i < 3; i++ {
			input := osMock.Mock.Open.GetInput(i)
			assert.Equal(t, fmt.Sprintf("test-name-%d", i), input.Name)
		}

		input := osMock.Mock.Open.GetInput(3)
		assert.Equal(t, "test-input-error", input.Name)
	})
}

func TestOsMockOpenFile(t *testing.T) {
	t.Helper()
	osMock := toolsmock.GetOsMock()

	t.Run("Testing OpenFile", func(t *testing.T) {
		var openFileReturn *os.File = &os.File{}

		osMock.Mock.OpenFile.AddReturnValue(&openFileReturn)
		osMock.Mock.OpenFile.AddReturnValue(&openFileReturn)
		osMock.Mock.OpenFile.AddReturnValue(&openFileReturn)

		for i := 0; i < 3; i++ {
			file, err := osMock.OpenFile(fmt.Sprintf("test-name-%d", i), os.O_RDONLY, 0644)
			assert.Equal(t, openFileReturn, file)
			assert.Nil(t, err)
		}

		file, err := osMock.OpenFile("test-input-error", os.O_RDONLY, 0644)
		assert.Nil(t, file)
		assert.Equal(t, fmt.Errorf("OpenFile general error"), err)

		assert.Equal(t, 4, osMock.Mock.OpenFile.HasBeenCalled())

		for i := 0; i < 3; i++ {
			input := osMock.Mock.OpenFile.GetInput(i)
			assert.Equal(t, fmt.Sprintf("test-name-%d", i), input.Name)
			assert.Equal(t, os.O_RDONLY, input.Flag)
			assert.Equal(t, os.FileMode(0644), input.Perm)
		}

		input := osMock.Mock.OpenFile.GetInput(3)
		assert.Equal(t, "test-input-error", input.Name)
		assert.Equal(t, os.O_RDONLY, input.Flag)
		assert.Equal(t, os.FileMode(0644), input.Perm)
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

func TestOsMockStat(t *testing.T) {
	t.Helper()
	osMock := toolsmock.GetOsMock()

	t.Run("Testing Stat", func(t *testing.T) {
		var statReturn fs.FileInfo = toolsmock.GetFileInfoMock()

		osMock.Mock.Stat.AddReturnValue(&statReturn)
		osMock.Mock.Stat.AddReturnValue(&statReturn)
		osMock.Mock.Stat.AddReturnValue(&statReturn)

		for i := 0; i < 3; i++ {
			content, err := osMock.Stat(fmt.Sprintf("test-name-%d", i))
			assert.Equal(t, statReturn, content)
			assert.Nil(t, err)
		}

		content, err := osMock.Stat("test-input-error")
		assert.Nil(t, content)
		assert.Equal(t, fmt.Errorf("Stat general error"), err)

		assert.Equal(t, 4, osMock.Mock.Stat.HasBeenCalled())
		for i := 0; i < 3; i++ {
			input := osMock.Mock.Stat.GetInput(i)
			assert.Equal(t, fmt.Sprintf("test-name-%d", i), input.Name)
		}

		input := osMock.Mock.Stat.GetInput(3)
		assert.Equal(t, "test-input-error", input.Name)
	})
}

func TestOsMockTempDir(t *testing.T) {
	t.Helper()
	osMock := toolsmock.GetOsMock()

	t.Run("Testing TempDir", func(t *testing.T) {
		tmpDir := "test-tmpdir"
		osMock.Mock.TempDir.AddReturnValue(&tmpDir)
		osMock.Mock.TempDir.AddReturnValue(&tmpDir)
		osMock.Mock.TempDir.AddReturnValue(&tmpDir)

		for i := 0; i < 3; i++ {
			dir := osMock.TempDir()
			assert.Equal(t, tmpDir, dir)
		}

		assert.Equal(t, 3, osMock.Mock.TempDir.HasBeenCalled())

		osMock.Mock.TempDir.Reset()
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}

			assert.Equal(t, 1, osMock.Mock.TempDir.HasBeenCalled())
		}()

		_ = osMock.TempDir()
	})
}

func TestOsMockUserHomeDir(t *testing.T) {
	t.Helper()
	osMock := toolsmock.GetOsMock()

	t.Run("Testing UserHomeDir", func(t *testing.T) {
		tmpDir := "test-tmpdir"
		osMock.Mock.UserHomeDir.AddReturnValue(&tmpDir)
		osMock.Mock.UserHomeDir.AddReturnValue(&tmpDir)
		osMock.Mock.UserHomeDir.AddReturnValue(&tmpDir)

		for i := 0; i < 3; i++ {
			dir, err := osMock.UserHomeDir()
			assert.Nil(t, err)
			assert.Equal(t, tmpDir, dir)
		}

		assert.Equal(t, 3, osMock.Mock.UserHomeDir.HasBeenCalled())

		dir, err := osMock.UserHomeDir()
		assert.Equal(t, "", dir)
		assert.Equal(t, fmt.Errorf("UserHomeDir general error"), err)

		assert.Equal(t, 4, osMock.Mock.UserHomeDir.HasBeenCalled())
	})
}
