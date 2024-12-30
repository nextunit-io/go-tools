package toolsmock_test

import (
	"os"
	"testing"
	"time"

	"github.com/nextunit-io/go-tools/toolsmock"
	"github.com/stretchr/testify/assert"
)

func TestFileInfoMockName(t *testing.T) {
	t.Helper()
	fsInfoMock := toolsmock.GetFileInfoMock()

	t.Run("Testing Name", func(t *testing.T) {
		fileName := "test-fileName"
		fsInfoMock.Mock.Name.AddReturnValue(&fileName)
		fsInfoMock.Mock.Name.AddReturnValue(&fileName)
		fsInfoMock.Mock.Name.AddReturnValue(&fileName)

		for i := 0; i < 3; i++ {
			dir := fsInfoMock.Name()
			assert.Equal(t, fileName, dir)
		}

		assert.Equal(t, 3, fsInfoMock.Mock.Name.HasBeenCalled())

		fsInfoMock.Mock.Name.Reset()
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}

			assert.Equal(t, 1, fsInfoMock.Mock.Name.HasBeenCalled())
		}()

		_ = fsInfoMock.Name()
	})
}

func TestFileInfoMockSize(t *testing.T) {
	t.Helper()
	fsInfoMock := toolsmock.GetFileInfoMock()

	t.Run("Testing Size", func(t *testing.T) {
		var size int64 = 123
		fsInfoMock.Mock.Size.AddReturnValue(&size)
		fsInfoMock.Mock.Size.AddReturnValue(&size)
		fsInfoMock.Mock.Size.AddReturnValue(&size)

		for i := 0; i < 3; i++ {
			dir := fsInfoMock.Size()
			assert.Equal(t, size, dir)
		}

		assert.Equal(t, 3, fsInfoMock.Mock.Size.HasBeenCalled())

		fsInfoMock.Mock.Size.Reset()
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}

			assert.Equal(t, 1, fsInfoMock.Mock.Size.HasBeenCalled())
		}()

		_ = fsInfoMock.Size()
	})
}

func TestFileInfoMockMode(t *testing.T) {
	t.Helper()
	fsInfoMock := toolsmock.GetFileInfoMock()

	t.Run("Testing Mode", func(t *testing.T) {
		var returnMode os.FileMode = 123
		fsInfoMock.Mock.Mode.AddReturnValue(&returnMode)
		fsInfoMock.Mock.Mode.AddReturnValue(&returnMode)
		fsInfoMock.Mock.Mode.AddReturnValue(&returnMode)

		for i := 0; i < 3; i++ {
			dir := fsInfoMock.Mode()
			assert.Equal(t, returnMode, dir)
		}

		assert.Equal(t, 3, fsInfoMock.Mock.Mode.HasBeenCalled())

		fsInfoMock.Mock.Mode.Reset()
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}

			assert.Equal(t, 1, fsInfoMock.Mock.Mode.HasBeenCalled())
		}()

		_ = fsInfoMock.Mode()
	})
}

func TestFileInfoMockModTime(t *testing.T) {
	t.Helper()
	fsInfoMock := toolsmock.GetFileInfoMock()

	t.Run("Testing ModTime", func(t *testing.T) {
		var returnModTime time.Time = time.Now()
		fsInfoMock.Mock.ModTime.AddReturnValue(&returnModTime)
		fsInfoMock.Mock.ModTime.AddReturnValue(&returnModTime)
		fsInfoMock.Mock.ModTime.AddReturnValue(&returnModTime)

		for i := 0; i < 3; i++ {
			dir := fsInfoMock.ModTime()
			assert.Equal(t, returnModTime, dir)
		}

		assert.Equal(t, 3, fsInfoMock.Mock.ModTime.HasBeenCalled())

		fsInfoMock.Mock.ModTime.Reset()
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}

			assert.Equal(t, 1, fsInfoMock.Mock.ModTime.HasBeenCalled())
		}()

		_ = fsInfoMock.ModTime()
	})
}

func TestFileInfoMockIsDir(t *testing.T) {
	t.Helper()
	fsInfoMock := toolsmock.GetFileInfoMock()

	t.Run("Testing IsDir", func(t *testing.T) {
		var returnIsDir bool = true
		fsInfoMock.Mock.IsDir.AddReturnValue(&returnIsDir)
		fsInfoMock.Mock.IsDir.AddReturnValue(&returnIsDir)
		fsInfoMock.Mock.IsDir.AddReturnValue(&returnIsDir)

		for i := 0; i < 3; i++ {
			dir := fsInfoMock.IsDir()
			assert.Equal(t, returnIsDir, dir)
		}

		assert.Equal(t, 3, fsInfoMock.Mock.IsDir.HasBeenCalled())

		fsInfoMock.Mock.IsDir.Reset()
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}

			assert.Equal(t, 1, fsInfoMock.Mock.IsDir.HasBeenCalled())
		}()

		_ = fsInfoMock.IsDir()
	})
}

func TestFileInfoMockSys(t *testing.T) {
	t.Helper()
	fsInfoMock := toolsmock.GetFileInfoMock()

	t.Run("Testing Sys", func(t *testing.T) {
		var returnSys any = "test"
		fsInfoMock.Mock.Sys.AddReturnValue(&returnSys)
		fsInfoMock.Mock.Sys.AddReturnValue(&returnSys)
		fsInfoMock.Mock.Sys.AddReturnValue(&returnSys)

		for i := 0; i < 3; i++ {
			dir := fsInfoMock.Sys()
			assert.Equal(t, returnSys, dir)
		}

		assert.Equal(t, 3, fsInfoMock.Mock.Sys.HasBeenCalled())

		fsInfoMock.Mock.Sys.Reset()
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}

			assert.Equal(t, 1, fsInfoMock.Mock.Sys.HasBeenCalled())
		}()

		_ = fsInfoMock.Sys()
	})
}
