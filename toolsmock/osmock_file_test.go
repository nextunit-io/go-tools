package toolsmock_test

import (
	"io/fs"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/nextunit-io/go-tools/toolsmock"
	"github.com/stretchr/testify/assert"
)

func TestFileMockChdir(t *testing.T) {
	t.Helper()
	fileMock := toolsmock.GetFileMock()

	t.Run("Testing Chdir", func(t *testing.T) {
		returnValue := false
		fileMock.Mock.Chdir.AddReturnValue(&returnValue)
		fileMock.Mock.Chdir.AddReturnValue(&returnValue)
		fileMock.Mock.Chdir.AddReturnValue(&returnValue)

		for i := 0; i < 3; i++ {
			err := fileMock.Chdir()
			assert.Nil(t, err)
		}

		err := fileMock.Chdir()
		assert.Equal(t, "Chdir general error", err.Error())

		assert.Equal(t, 4, fileMock.Mock.Chdir.HasBeenCalled())
	})
}
func TestFileMockChmod(t *testing.T) {
	t.Helper()
	fileMock := toolsmock.GetFileMock()

	t.Run("Testing Chmod", func(t *testing.T) {
		returnValue := false
		fileMock.Mock.Chmod.AddReturnValue(&returnValue)
		fileMock.Mock.Chmod.AddReturnValue(&returnValue)
		fileMock.Mock.Chmod.AddReturnValue(&returnValue)

		for i := 0; i < 3; i++ {
			err := fileMock.Chmod(fs.ModeDevice)
			assert.Nil(t, err)
		}

		err := fileMock.Chmod(fs.ModeDevice)
		assert.Equal(t, "Chmod general error", err.Error())

		assert.Equal(t, 4, fileMock.Mock.Chmod.HasBeenCalled())

		for i := 0; i < 4; i++ {
			input := fileMock.Mock.Chmod.GetInput(i)
			assert.Equal(t, fs.ModeDevice, input.Mode)
		}
	})
}

func TestFileMockChown(t *testing.T) {
	t.Helper()
	fileMock := toolsmock.GetFileMock()

	t.Run("Testing Chown", func(t *testing.T) {
		returnValue := false
		fileMock.Mock.Chown.AddReturnValue(&returnValue)
		fileMock.Mock.Chown.AddReturnValue(&returnValue)
		fileMock.Mock.Chown.AddReturnValue(&returnValue)

		for i := 0; i < 3; i++ {
			err := fileMock.Chown(i, i+1)
			assert.Nil(t, err)
		}

		err := fileMock.Chown(3, 4)
		assert.Equal(t, "Chown general error", err.Error())

		assert.Equal(t, 4, fileMock.Mock.Chown.HasBeenCalled())

		for i := 0; i < 4; i++ {
			input := fileMock.Mock.Chown.GetInput(i)
			assert.Equal(t, i, input.UID)
			assert.Equal(t, i+1, input.GID)
		}
	})
}

func TestFileMockClose(t *testing.T) {
	t.Helper()
	fileMock := toolsmock.GetFileMock()

	t.Run("Testing Close", func(t *testing.T) {
		returnValue := false
		fileMock.Mock.Close.AddReturnValue(&returnValue)
		fileMock.Mock.Close.AddReturnValue(&returnValue)
		fileMock.Mock.Close.AddReturnValue(&returnValue)

		for i := 0; i < 3; i++ {
			err := fileMock.Close()
			assert.Nil(t, err)
		}

		err := fileMock.Close()
		assert.Equal(t, "Close general error", err.Error())

		assert.Equal(t, 4, fileMock.Mock.Close.HasBeenCalled())
	})
}

func TestFileMockFd(t *testing.T) {
	t.Helper()
	fileMock := toolsmock.GetFileMock()

	t.Run("Testing Fd", func(t *testing.T) {
		returnValue := uintptr(0)
		fileMock.Mock.Fd.AddReturnValue(&returnValue)
		fileMock.Mock.Fd.AddReturnValue(&returnValue)
		fileMock.Mock.Fd.AddReturnValue(&returnValue)

		for i := 0; i < 3; i++ {
			fd := fileMock.Fd()
			assert.Equal(t, returnValue, fd)
		}

		fileMock.Mock.Fd.Reset()
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}

			assert.Equal(t, 1, fileMock.Mock.Fd.HasBeenCalled())
		}()

		_ = fileMock.Fd()
	})
}

func TestFileMockName(t *testing.T) {
	t.Helper()
	fileMock := toolsmock.GetFileMock()

	t.Run("Testing Name", func(t *testing.T) {
		returnValue := "testfile"
		fileMock.Mock.Name.AddReturnValue(&returnValue)
		fileMock.Mock.Name.AddReturnValue(&returnValue)
		fileMock.Mock.Name.AddReturnValue(&returnValue)

		for i := 0; i < 3; i++ {
			name := fileMock.Name()
			assert.Equal(t, returnValue, name)
		}

		fileMock.Mock.Name.Reset()
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}

			assert.Equal(t, 1, fileMock.Mock.Name.HasBeenCalled())
		}()

		_ = fileMock.Name()
	})
}

func TestFileMockRead(t *testing.T) {
	t.Helper()
	fileMock := toolsmock.GetFileMock()

	t.Run("Testing Read", func(t *testing.T) {
		returnValue := 5
		fileMock.Mock.Read.AddReturnValue(&returnValue)
		fileMock.Mock.Read.AddReturnValue(&returnValue)
		fileMock.Mock.Read.AddReturnValue(&returnValue)

		input := make([]byte, 10)
		for i := 0; i < 3; i++ {
			n, err := fileMock.Read(input)
			assert.Nil(t, err)
			assert.Equal(t, returnValue, n)
		}

		n, err := fileMock.Read(input)
		assert.Equal(t, 0, n)
		assert.Equal(t, "Read general error", err.Error())

		assert.Equal(t, 4, fileMock.Mock.Read.HasBeenCalled())

		for i := 0; i < 4; i++ {
			inputResult := fileMock.Mock.Read.GetInput(i)
			assert.Equal(t, input, inputResult.B)
		}
	})
}

func TestFileMockReadAt(t *testing.T) {
	t.Helper()
	fileMock := toolsmock.GetFileMock()

	t.Run("Testing ReadAt", func(t *testing.T) {
		returnValue := 5
		fileMock.Mock.ReadAt.AddReturnValue(&returnValue)
		fileMock.Mock.ReadAt.AddReturnValue(&returnValue)
		fileMock.Mock.ReadAt.AddReturnValue(&returnValue)

		input := make([]byte, 10)
		for i := 0; i < 3; i++ {
			n, err := fileMock.ReadAt(input, int64(i))
			assert.Nil(t, err)
			assert.Equal(t, returnValue, n)
		}

		n, err := fileMock.ReadAt(input, int64(3))
		assert.Equal(t, 0, n)
		assert.Equal(t, "ReadAt general error", err.Error())

		assert.Equal(t, 4, fileMock.Mock.ReadAt.HasBeenCalled())

		for i := 0; i < 4; i++ {
			inputResult := fileMock.Mock.ReadAt.GetInput(i)
			assert.Equal(t, input, inputResult.B)
		}
	})
}

func TestFileMockReadDir(t *testing.T) {
	t.Helper()
	fileMock := toolsmock.GetFileMock()

	t.Run("Testing ReadDir", func(t *testing.T) {
		returnValue := []fs.DirEntry{}
		fileMock.Mock.ReadDir.AddReturnValue(&returnValue)
		fileMock.Mock.ReadDir.AddReturnValue(&returnValue)
		fileMock.Mock.ReadDir.AddReturnValue(&returnValue)

		input := 10
		for i := 0; i < 3; i++ {
			entries, err := fileMock.ReadDir(input)
			assert.Nil(t, err)
			assert.Equal(t, returnValue, entries)
		}

		entry, err := fileMock.ReadDir(input)
		assert.Equal(t, 0, len(entry))
		assert.Equal(t, "ReadDir general error", err.Error())

		assert.Equal(t, 4, fileMock.Mock.ReadDir.HasBeenCalled())

		for i := 0; i < 4; i++ {
			inputResult := fileMock.Mock.ReadDir.GetInput(i)
			assert.Equal(t, input, inputResult.N)
		}
	})
}

func TestFileMockReadFrom(t *testing.T) {
	t.Helper()
	fileMock := toolsmock.GetFileMock()

	t.Run("Testing ReadFrom", func(t *testing.T) {
		returnValue := int64(5)
		fileMock.Mock.ReadFrom.AddReturnValue(&returnValue)
		fileMock.Mock.ReadFrom.AddReturnValue(&returnValue)
		fileMock.Mock.ReadFrom.AddReturnValue(&returnValue)

		input := strings.NewReader("test")
		for i := 0; i < 3; i++ {
			n, err := fileMock.ReadFrom(input)
			assert.Nil(t, err)
			assert.Equal(t, returnValue, n)
		}

		n, err := fileMock.ReadFrom(input)
		assert.Equal(t, int64(0), n)
		assert.Equal(t, "ReadFrom general error", err.Error())

		assert.Equal(t, 4, fileMock.Mock.ReadFrom.HasBeenCalled())

		for i := 0; i < 4; i++ {
			inputResult := fileMock.Mock.ReadFrom.GetInput(i)
			assert.Equal(t, input, inputResult.R)
		}
	})
}

func TestFileMockReaddir(t *testing.T) {
	t.Helper()
	fileMock := toolsmock.GetFileMock()

	t.Run("Testing Readdir", func(t *testing.T) {
		returnValue := []fs.FileInfo{}
		fileMock.Mock.Readdir.AddReturnValue(&returnValue)
		fileMock.Mock.Readdir.AddReturnValue(&returnValue)
		fileMock.Mock.Readdir.AddReturnValue(&returnValue)

		input := 10
		for i := 0; i < 3; i++ {
			infos, err := fileMock.Readdir(input)
			assert.Nil(t, err)
			assert.Equal(t, returnValue, infos)
		}

		infos, err := fileMock.Readdir(input)
		assert.Equal(t, 0, len(infos))
		assert.Equal(t, "Readdir general error", err.Error())

		assert.Equal(t, 4, fileMock.Mock.Readdir.HasBeenCalled())

		for i := 0; i < 4; i++ {
			inputResult := fileMock.Mock.Readdir.GetInput(i)
			assert.Equal(t, input, inputResult.N)
		}
	})
}

func TestFileMockReaddirnames(t *testing.T) {
	t.Helper()
	fileMock := toolsmock.GetFileMock()

	t.Run("Testing Readdirnames", func(t *testing.T) {
		returnValue := []string{}
		fileMock.Mock.Readdirnames.AddReturnValue(&returnValue)
		fileMock.Mock.Readdirnames.AddReturnValue(&returnValue)
		fileMock.Mock.Readdirnames.AddReturnValue(&returnValue)

		input := 10
		for i := 0; i < 3; i++ {
			names, err := fileMock.Readdirnames(input)
			assert.Nil(t, err)
			assert.Equal(t, returnValue, names)
		}

		names, err := fileMock.Readdirnames(input)
		assert.Equal(t, 0, len(names))
		assert.Equal(t, "Readdirnames general error", err.Error())

		assert.Equal(t, 4, fileMock.Mock.Readdirnames.HasBeenCalled())
	})
}

func TestFileMockSeek(t *testing.T) {
	t.Helper()
	fileMock := toolsmock.GetFileMock()

	t.Run("Testing Seek", func(t *testing.T) {
		returnValue := int64(5)
		fileMock.Mock.Seek.AddReturnValue(&returnValue)
		fileMock.Mock.Seek.AddReturnValue(&returnValue)
		fileMock.Mock.Seek.AddReturnValue(&returnValue)

		input1 := 10
		input2 := 0
		for i := 0; i < 3; i++ {
			offset, err := fileMock.Seek(int64(input1), int(input2))
			assert.Nil(t, err)
			assert.Equal(t, returnValue, offset)
		}

		offset, err := fileMock.Seek(int64(input1), int(input2))
		assert.Equal(t, int64(0), offset)
		assert.Equal(t, "Seek general error", err.Error())

		assert.Equal(t, 4, fileMock.Mock.Seek.HasBeenCalled())
	})
}

func TestFileMockSetDeadline(t *testing.T) {
	t.Helper()
	fileMock := toolsmock.GetFileMock()

	t.Run("Testing SetDeadline", func(t *testing.T) {
		returnValue := false
		fileMock.Mock.SetDeadline.AddReturnValue(&returnValue)
		fileMock.Mock.SetDeadline.AddReturnValue(&returnValue)
		fileMock.Mock.SetDeadline.AddReturnValue(&returnValue)

		for i := 0; i < 3; i++ {
			err := fileMock.SetDeadline(time.Now())
			assert.Nil(t, err)
		}

		err := fileMock.SetDeadline(time.Now())
		assert.Equal(t, "SetDeadline general error", err.Error())

		assert.Equal(t, 4, fileMock.Mock.SetDeadline.HasBeenCalled())
	})
}

func TestFileMockSetReadDeadline(t *testing.T) {
	t.Helper()
	fileMock := toolsmock.GetFileMock()

	t.Run("Testing SetReadDeadline", func(t *testing.T) {
		returnValue := false
		fileMock.Mock.SetReadDeadline.AddReturnValue(&returnValue)
		fileMock.Mock.SetReadDeadline.AddReturnValue(&returnValue)
		fileMock.Mock.SetReadDeadline.AddReturnValue(&returnValue)

		for i := 0; i < 3; i++ {
			err := fileMock.SetReadDeadline(time.Now())
			assert.Nil(t, err)
		}

		err := fileMock.SetReadDeadline(time.Now())
		assert.Equal(t, "SetReadDeadline general error", err.Error())

		assert.Equal(t, 4, fileMock.Mock.SetReadDeadline.HasBeenCalled())
	})
}

func TestFileMockSetWriteDeadline(t *testing.T) {
	t.Helper()
	fileMock := toolsmock.GetFileMock()

	t.Run("Testing SetWriteDeadline", func(t *testing.T) {
		returnValue := false
		fileMock.Mock.SetWriteDeadline.AddReturnValue(&returnValue)
		fileMock.Mock.SetWriteDeadline.AddReturnValue(&returnValue)
		fileMock.Mock.SetWriteDeadline.AddReturnValue(&returnValue)

		for i := 0; i < 3; i++ {
			err := fileMock.SetWriteDeadline(time.Now())
			assert.Nil(t, err)
		}

		err := fileMock.SetWriteDeadline(time.Now())
		assert.Equal(t, "SetWriteDeadline general error", err.Error())

		assert.Equal(t, 4, fileMock.Mock.SetWriteDeadline.HasBeenCalled())
	})
}

func TestFileMockStat(t *testing.T) {
	t.Helper()
	fileMock := toolsmock.GetFileMock()

	t.Run("Testing Stat", func(t *testing.T) {
		returnValue := fs.FileInfo(nil)
		fileMock.Mock.Stat.AddReturnValue(&returnValue)
		fileMock.Mock.Stat.AddReturnValue(&returnValue)
		fileMock.Mock.Stat.AddReturnValue(&returnValue)

		for i := 0; i < 3; i++ {
			info, err := fileMock.Stat()
			assert.Nil(t, err)
			assert.Equal(t, returnValue, info)
		}

		info, err := fileMock.Stat()
		assert.Nil(t, info)
		assert.Equal(t, "Stat general error", err.Error())

		assert.Equal(t, 4, fileMock.Mock.Stat.HasBeenCalled())
	})
}

func TestFileMockSync(t *testing.T) {
	t.Helper()
	fileMock := toolsmock.GetFileMock()

	t.Run("Testing Sync", func(t *testing.T) {
		returnValue := false
		fileMock.Mock.Sync.AddReturnValue(&returnValue)
		fileMock.Mock.Sync.AddReturnValue(&returnValue)
		fileMock.Mock.Sync.AddReturnValue(&returnValue)

		for i := 0; i < 3; i++ {
			err := fileMock.Sync()
			assert.Nil(t, err)
		}

		err := fileMock.Sync()
		assert.Equal(t, "Sync general error", err.Error())

		assert.Equal(t, 4, fileMock.Mock.Sync.HasBeenCalled())
	})
}

func TestFileMockSyscallConn(t *testing.T) {
	t.Helper()
	fileMock := toolsmock.GetFileMock()

	t.Run("Testing SyscallConn", func(t *testing.T) {
		returnValue := syscall.RawConn(nil)
		fileMock.Mock.SyscallConn.AddReturnValue(&returnValue)
		fileMock.Mock.SyscallConn.AddReturnValue(&returnValue)
		fileMock.Mock.SyscallConn.AddReturnValue(&returnValue)

		for i := 0; i < 3; i++ {
			conn, err := fileMock.SyscallConn()
			assert.Nil(t, err)
			assert.Equal(t, returnValue, conn)
		}

		conn, err := fileMock.SyscallConn()
		assert.Nil(t, conn)
		assert.Equal(t, "SyscallConn general error", err.Error())

		assert.Equal(t, 4, fileMock.Mock.SyscallConn.HasBeenCalled())
	})
}

func TestFileMockTruncate(t *testing.T) {
	t.Helper()
	fileMock := toolsmock.GetFileMock()

	t.Run("Testing Truncate", func(t *testing.T) {
		returnValue := false
		fileMock.Mock.Truncate.AddReturnValue(&returnValue)
		fileMock.Mock.Truncate.AddReturnValue(&returnValue)
		fileMock.Mock.Truncate.AddReturnValue(&returnValue)

		for i := 0; i < 3; i++ {
			err := fileMock.Truncate(10)
			assert.Nil(t, err)
		}

		err := fileMock.Truncate(10)
		assert.Equal(t, "Truncate general error", err.Error())

		assert.Equal(t, 4, fileMock.Mock.Truncate.HasBeenCalled())
	})
}

func TestFileMockWrite(t *testing.T) {
	t.Helper()
	fileMock := toolsmock.GetFileMock()

	t.Run("Testing Write", func(t *testing.T) {
		returnValue := 5
		fileMock.Mock.Write.AddReturnValue(&returnValue)
		fileMock.Mock.Write.AddReturnValue(&returnValue)
		fileMock.Mock.Write.AddReturnValue(&returnValue)

		input := make([]byte, 10)
		for i := 0; i < 3; i++ {
			n, err := fileMock.Write(input)
			assert.Nil(t, err)
			assert.Equal(t, returnValue, n)
		}

		n, err := fileMock.Write(input)
		assert.Equal(t, 0, n)
		assert.Equal(t, "Write general error", err.Error())

		assert.Equal(t, 4, fileMock.Mock.Write.HasBeenCalled())

		for i := 0; i < 4; i++ {
			inputResult := fileMock.Mock.Write.GetInput(i)
			assert.Equal(t, input, inputResult.B)
		}
	})
}

func TestFileMockWriteAt(t *testing.T) {
	t.Helper()
	fileMock := toolsmock.GetFileMock()

	t.Run("Testing WriteAt", func(t *testing.T) {
		returnValue := 5
		fileMock.Mock.WriteAt.AddReturnValue(&returnValue)
		fileMock.Mock.WriteAt.AddReturnValue(&returnValue)
		fileMock.Mock.WriteAt.AddReturnValue(&returnValue)

		input1 := make([]byte, 10)
		for i := 0; i < 3; i++ {
			n, err := fileMock.WriteAt(input1, int64(i))
			assert.Nil(t, err)
			assert.Equal(t, returnValue, n)
		}

		n, err := fileMock.WriteAt(input1, 3)
		assert.Equal(t, 0, n)
		assert.Equal(t, "WriteAt general error", err.Error())

		assert.Equal(t, 4, fileMock.Mock.WriteAt.HasBeenCalled())

		for i := 0; i < 4; i++ {
			inputResult := fileMock.Mock.WriteAt.GetInput(i)
			assert.Equal(t, input1, inputResult.B)
			assert.Equal(t, int64(i), inputResult.Off)
		}
	})
}

func TestFileMockWriteString(t *testing.T) {
	t.Helper()
	fileMock := toolsmock.GetFileMock()

	t.Run("Testing WriteString", func(t *testing.T) {
		returnValue := 5
		fileMock.Mock.WriteString.AddReturnValue(&returnValue)
		fileMock.Mock.WriteString.AddReturnValue(&returnValue)
		fileMock.Mock.WriteString.AddReturnValue(&returnValue)

		for i := 0; i < 3; i++ {
			n, err := fileMock.WriteString("test")
			assert.Nil(t, err)
			assert.Equal(t, returnValue, n)
		}

		n, err := fileMock.WriteString("test-error")
		assert.Equal(t, 0, n)
		assert.Equal(t, "WriteString general error", err.Error())

		assert.Equal(t, 4, fileMock.Mock.WriteString.HasBeenCalled())

		for i := 0; i < 3; i++ {
			inputResult := fileMock.Mock.WriteString.GetInput(i)
			assert.Equal(t, "test", inputResult.S)
		}

		inputResult := fileMock.Mock.WriteString.GetInput(3)
		assert.Equal(t, "test-error", inputResult.S)
	})
}

func TestFileMockWriteTo(t *testing.T) {
	t.Helper()
	fileMock := toolsmock.GetFileMock()

	t.Run("Testing WriteTo", func(t *testing.T) {
		returnValue := int64(5)
		fileMock.Mock.WriteTo.AddReturnValue(&returnValue)
		fileMock.Mock.WriteTo.AddReturnValue(&returnValue)
		fileMock.Mock.WriteTo.AddReturnValue(&returnValue)

		input := new(strings.Builder)
		for i := 0; i < 3; i++ {
			n, err := fileMock.WriteTo(input)
			assert.Nil(t, err)
			assert.Equal(t, returnValue, n)
		}

		n, err := fileMock.WriteTo(input)
		assert.Equal(t, int64(0), n)
		assert.Equal(t, "WriteTo general error", err.Error())

		assert.Equal(t, 4, fileMock.Mock.WriteTo.HasBeenCalled())
	})
}
