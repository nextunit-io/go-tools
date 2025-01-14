package toolsmock_test

import (
	"os"
	"testing"

	"github.com/nextunit-io/go-tools/toolsmock"
	"github.com/stretchr/testify/assert"
)

func TestDirEntryMockName(t *testing.T) {
	t.Helper()
	mock := toolsmock.GetDirEntryMock()

	t.Run("Testing Name", func(t *testing.T) {
		fileName := "test-fileName"
		mock.Mock.Name.AddReturnValue(&fileName)
		mock.Mock.Name.AddReturnValue(&fileName)
		mock.Mock.Name.AddReturnValue(&fileName)

		for i := 0; i < 3; i++ {
			dir := mock.Name()
			assert.Equal(t, fileName, dir)
		}

		assert.Equal(t, 3, mock.Mock.Name.HasBeenCalled())

		mock.Mock.Name.Reset()
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}

			assert.Equal(t, 1, mock.Mock.Name.HasBeenCalled())
		}()

		_ = mock.Name()
	})
}

func TestDirEntryMockIsDir(t *testing.T) {
	t.Helper()
	mock := toolsmock.GetDirEntryMock()

	t.Run("Testing IsDir", func(t *testing.T) {
		isDir := true
		mock.Mock.IsDir.AddReturnValue(&isDir)
		mock.Mock.IsDir.AddReturnValue(&isDir)
		mock.Mock.IsDir.AddReturnValue(&isDir)

		for i := 0; i < 3; i++ {
			dir := mock.IsDir()
			assert.Equal(t, isDir, dir)
		}

		assert.Equal(t, 3, mock.Mock.IsDir.HasBeenCalled())

		mock.Mock.IsDir.Reset()
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}

			assert.Equal(t, 1, mock.Mock.IsDir.HasBeenCalled())
		}()

		_ = mock.IsDir()
	})
}

func TestDirEntryMockType(t *testing.T) {
	t.Helper()
	mock := toolsmock.GetDirEntryMock()

	t.Run("Testing Type", func(t *testing.T) {
		fileMode := os.FileMode(0755)
		mock.Mock.Type.AddReturnValue(&fileMode)
		mock.Mock.Type.AddReturnValue(&fileMode)
		mock.Mock.Type.AddReturnValue(&fileMode)

		for i := 0; i < 3; i++ {
			mode := mock.Type()
			assert.Equal(t, fileMode, mode)
		}

		assert.Equal(t, 3, mock.Mock.Type.HasBeenCalled())

		mock.Mock.Type.Reset()
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}

			assert.Equal(t, 1, mock.Mock.Type.HasBeenCalled())
		}()

		_ = mock.Type()
	})
}

func TestDirEntryMockInfo(t *testing.T) {
	t.Helper()
	mock := toolsmock.GetDirEntryMock()

	t.Run("Testing Info", func(t *testing.T) {
		var fileInfo os.FileInfo = toolsmock.GetFileInfoMock()
		mock.Mock.Info.AddReturnValue(&fileInfo)
		mock.Mock.Info.AddReturnValue(&fileInfo)
		mock.Mock.Info.AddReturnValue(&fileInfo)

		for i := 0; i < 3; i++ {
			info, err := mock.Info()
			assert.Nil(t, err)
			assert.Equal(t, fileInfo, info)
		}

		assert.Equal(t, 3, mock.Mock.Info.HasBeenCalled())

		mock.Mock.Info.Reset()

		_, _ = mock.Info()
	})
}
