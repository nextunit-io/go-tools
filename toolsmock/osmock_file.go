package toolsmock

import (
	"fmt"
	"io"
	"os"
	"syscall"
	"time"

	gomock "github.com/nextunit-io/go-mock"
)

type fileMockStruct struct {
	Chdir *gomock.ToolMock[
		interface{},
		bool,
	]
	Chmod *gomock.ToolMock[
		struct {
			Mode os.FileMode
		},
		bool,
	]
	Chown *gomock.ToolMock[
		struct {
			UID int
			GID int
		},
		bool,
	]
	Close *gomock.ToolMock[
		interface{},
		bool,
	]
	Fd *gomock.ToolMock[
		interface{},
		uintptr,
	]
	IsDir *gomock.ToolMock[
		interface{},
		bool,
	]
	ModTime *gomock.ToolMock[
		interface{},
		time.Time,
	]
	Mode *gomock.ToolMock[
		interface{},
		os.FileMode,
	]
	Name *gomock.ToolMock[
		interface{},
		string,
	]
	Read *gomock.ToolMock[
		struct {
			B []byte
		},
		int,
	]
	ReadAt *gomock.ToolMock[
		struct {
			B   []byte
			Off int64
		},
		int,
	]
	ReadDir *gomock.ToolMock[
		struct {
			N int
		},
		[]os.DirEntry,
	]
	ReadFrom *gomock.ToolMock[
		struct {
			R io.Reader
		},
		int64,
	]
	Readdir *gomock.ToolMock[
		struct {
			N int
		},
		[]os.FileInfo,
	]
	Readdirnames *gomock.ToolMock[
		struct {
			N int
		},
		[]string,
	]
	Seek *gomock.ToolMock[
		struct {
			Offset int64
			Whence int
		},
		int64,
	]
	SetDeadline *gomock.ToolMock[
		struct {
			T time.Time
		},
		bool,
	]
	SetReadDeadline *gomock.ToolMock[
		struct {
			T time.Time
		},
		bool,
	]
	SetWriteDeadline *gomock.ToolMock[
		struct {
			T time.Time
		},
		bool,
	]
	Size *gomock.ToolMock[
		interface{},
		int64,
	]
	Stat *gomock.ToolMock[
		interface{},
		os.FileInfo,
	]
	Sync *gomock.ToolMock[
		interface{},
		bool,
	]
	Sys *gomock.ToolMock[
		interface{},
		any,
	]
	SyscallConn *gomock.ToolMock[
		interface{},
		syscall.RawConn,
	]
	Truncate *gomock.ToolMock[
		struct {
			Size int64
		},
		bool,
	]
	Write *gomock.ToolMock[
		struct {
			B []byte
		},
		int,
	]
	WriteAt *gomock.ToolMock[
		struct {
			B   []byte
			Off int64
		},
		int,
	]
	WriteString *gomock.ToolMock[
		struct {
			S string
		},
		int,
	]
	WriteTo *gomock.ToolMock[
		struct {
			W io.Writer
		},
		int64,
	]
}

type FileMock struct {
	Mock fileMockStruct
}

func GetFileMock() *FileMock {
	return &FileMock{
		Mock: fileMockStruct{
			Chdir: gomock.GetMock[
				interface{},
				bool,
			](fmt.Errorf("Chdir general error")),
			Chmod: gomock.GetMock[
				struct {
					Mode os.FileMode
				},
				bool,
			](fmt.Errorf("Chmod general error")),
			Chown: gomock.GetMock[
				struct {
					UID int
					GID int
				},
				bool,
			](fmt.Errorf("Chown general error")),
			Close: gomock.GetMock[
				interface{},
				bool,
			](fmt.Errorf("Close general error")),
			Fd: gomock.GetMock[
				interface{},
				uintptr,
			](fmt.Errorf("Fd general error")),
			IsDir: gomock.GetMock[
				interface{},
				bool,
			](fmt.Errorf("IsDir general error")),
			ModTime: gomock.GetMock[
				interface{},
				time.Time,
			](fmt.Errorf("ModTime general error")),
			Mode: gomock.GetMock[
				interface{},
				os.FileMode,
			](fmt.Errorf("Mode general error")),
			Name: gomock.GetMock[
				interface{},
				string,
			](fmt.Errorf("Name general error")),
			Read: gomock.GetMock[
				struct {
					B []byte
				},
				int,
			](fmt.Errorf("Read general error")),
			ReadAt: gomock.GetMock[
				struct {
					B   []byte
					Off int64
				},
				int,
			](fmt.Errorf("ReadAt general error")),
			ReadDir: gomock.GetMock[
				struct {
					N int
				},
				[]os.DirEntry,
			](fmt.Errorf("ReadDir general error")),
			ReadFrom: gomock.GetMock[
				struct {
					R io.Reader
				},
				int64,
			](fmt.Errorf("ReadFrom general error")),
			Readdir: gomock.GetMock[
				struct {
					N int
				},
				[]os.FileInfo,
			](fmt.Errorf("Readdir general error")),
			Readdirnames: gomock.GetMock[
				struct {
					N int
				},
				[]string,
			](fmt.Errorf("Readdirnames general error")),
			Seek: gomock.GetMock[
				struct {
					Offset int64
					Whence int
				},
				int64,
			](fmt.Errorf("Seek general error")),
			SetDeadline: gomock.GetMock[
				struct {
					T time.Time
				},
				bool,
			](fmt.Errorf("SetDeadline general error")),
			SetReadDeadline: gomock.GetMock[
				struct {
					T time.Time
				},
				bool,
			](fmt.Errorf("SetReadDeadline general error")),
			SetWriteDeadline: gomock.GetMock[
				struct {
					T time.Time
				},
				bool,
			](fmt.Errorf("SetWriteDeadline general error")),
			Size: gomock.GetMock[
				interface{},
				int64,
			](fmt.Errorf("Size general error")),
			Stat: gomock.GetMock[
				interface{},
				os.FileInfo,
			](fmt.Errorf("Stat general error")),
			Sync: gomock.GetMock[
				interface{},
				bool,
			](fmt.Errorf("Sync general error")),
			Sys: gomock.GetMock[
				interface{},
				any,
			](fmt.Errorf("Sys general error")),
			SyscallConn: gomock.GetMock[
				interface{},
				syscall.RawConn,
			](fmt.Errorf("SyscallConn general error")),
			Truncate: gomock.GetMock[
				struct {
					Size int64
				},
				bool,
			](fmt.Errorf("Truncate general error")),
			Write: gomock.GetMock[
				struct {
					B []byte
				},
				int,
			](fmt.Errorf("Write general error")),
			WriteAt: gomock.GetMock[
				struct {
					B   []byte
					Off int64
				},
				int,
			](fmt.Errorf("WriteAt general error")),
			WriteString: gomock.GetMock[
				struct {
					S string
				},
				int,
			](fmt.Errorf("WriteString general error")),
			WriteTo: gomock.GetMock[
				struct {
					W io.Writer
				},
				int64,
			](fmt.Errorf("WriteTo general error")),
		},
	}
}

func (fileMock *FileMock) Chdir() error {
	fileMock.Mock.Chdir.AddInput(struct{}{})

	result, err := fileMock.Mock.Chdir.GetNextResult()
	if err != nil {
		return err
	}

	if *result {
		return fmt.Errorf("Error requested")
	}

	return nil
}

func (fileMock *FileMock) Chmod(mode os.FileMode) error {
	fileMock.Mock.Chmod.AddInput(struct {
		Mode os.FileMode
	}{
		mode,
	})

	result, err := fileMock.Mock.Chmod.GetNextResult()
	if err != nil {
		return err
	}

	if *result {
		return fmt.Errorf("Error requested")
	}

	return nil
}

func (fileMock *FileMock) Chown(uid int, gid int) error {
	fileMock.Mock.Chown.AddInput(struct {
		UID int
		GID int
	}{
		uid,
		gid,
	})

	result, err := fileMock.Mock.Chown.GetNextResult()
	if err != nil {
		return err
	}

	if *result {
		return fmt.Errorf("Error requested")
	}

	return nil
}

func (fileMock *FileMock) Close() error {
	fileMock.Mock.Close.AddInput(struct{}{})

	result, err := fileMock.Mock.Close.GetNextResult()
	if err != nil {
		return err
	}

	if *result {
		return fmt.Errorf("Error requested")
	}

	return nil
}

func (fileMock *FileMock) Fd() uintptr {
	fileMock.Mock.Fd.AddInput(struct{}{})

	result, err := fileMock.Mock.Fd.GetNextResult()
	if err != nil {
		panic(err)
	}

	return *result
}

func (fileMock *FileMock) Name() string {
	fileMock.Mock.Name.AddInput(struct{}{})

	result, _ := fileMock.Mock.Name.GetNextResult()
	return *result
}

func (fileMock *FileMock) Read(b []byte) (n int, err error) {
	fileMock.Mock.Read.AddInput(struct {
		B []byte
	}{
		B: b,
	})

	result, err := fileMock.Mock.Read.GetNextResult()
	if err != nil {
		return 0, err
	}

	return *result, nil
}

func (fileMock *FileMock) ReadAt(b []byte, off int64) (n int, err error) {
	fileMock.Mock.ReadAt.AddInput(struct {
		B   []byte
		Off int64
	}{
		B:   b,
		Off: off,
	})

	result, err := fileMock.Mock.ReadAt.GetNextResult()
	if err != nil {
		return 0, err
	}

	return *result, nil
}

func (fileMock *FileMock) ReadDir(n int) ([]os.DirEntry, error) {
	fileMock.Mock.ReadDir.AddInput(struct {
		N int
	}{
		N: n,
	})

	result, err := fileMock.Mock.ReadDir.GetNextResult()
	if err != nil {
		return nil, err
	}

	return *result, nil
}

func (fileMock *FileMock) ReadFrom(r io.Reader) (n int64, err error) {
	fileMock.Mock.ReadFrom.AddInput(struct {
		R io.Reader
	}{
		R: r,
	})

	result, err := fileMock.Mock.ReadFrom.GetNextResult()
	if err != nil {
		return 0, err
	}

	return *result, nil
}

func (fileMock *FileMock) Readdir(n int) ([]os.FileInfo, error) {
	fileMock.Mock.Readdir.AddInput(struct {
		N int
	}{
		N: n,
	})

	result, err := fileMock.Mock.Readdir.GetNextResult()
	if err != nil {
		return nil, err
	}

	return *result, nil
}

func (fileMock *FileMock) Readdirnames(n int) ([]string, error) {
	fileMock.Mock.Readdirnames.AddInput(struct {
		N int
	}{
		N: n,
	})

	result, err := fileMock.Mock.Readdirnames.GetNextResult()
	if err != nil {
		return nil, err
	}

	return *result, nil
}

func (fileMock *FileMock) Seek(offset int64, whence int) (ret int64, err error) {
	fileMock.Mock.Seek.AddInput(struct {
		Offset int64
		Whence int
	}{
		Offset: offset,
		Whence: whence,
	})

	result, err := fileMock.Mock.Seek.GetNextResult()
	if err != nil {
		return 0, err
	}

	return *result, nil
}

func (fileMock *FileMock) SetDeadline(t time.Time) error {
	fileMock.Mock.SetDeadline.AddInput(struct {
		T time.Time
	}{
		T: t,
	})

	result, err := fileMock.Mock.SetDeadline.GetNextResult()
	if err != nil {
		return err
	}

	if *result {
		return fmt.Errorf("Error requested")
	}

	return nil
}

func (fileMock *FileMock) SetReadDeadline(t time.Time) error {
	fileMock.Mock.SetReadDeadline.AddInput(struct {
		T time.Time
	}{
		T: t,
	})

	result, err := fileMock.Mock.SetReadDeadline.GetNextResult()
	if err != nil {
		return err
	}

	if *result {
		return fmt.Errorf("Error requested")
	}

	return nil
}

func (fileMock *FileMock) SetWriteDeadline(t time.Time) error {
	fileMock.Mock.SetWriteDeadline.AddInput(struct {
		T time.Time
	}{
		T: t,
	})

	result, err := fileMock.Mock.SetWriteDeadline.GetNextResult()
	if err != nil {
		return err
	}

	if *result {
		return fmt.Errorf("Error requested")
	}

	return nil
}

func (fileMock *FileMock) Stat() (os.FileInfo, error) {
	fileMock.Mock.Stat.AddInput(struct{}{})

	result, err := fileMock.Mock.Stat.GetNextResult()
	if err != nil {
		return nil, err
	}

	return *result, nil
}

func (fileMock *FileMock) Sync() error {
	fileMock.Mock.Sync.AddInput(struct{}{})

	result, err := fileMock.Mock.Sync.GetNextResult()
	if err != nil {
		return err
	}

	if *result {
		return fmt.Errorf("Error requested")
	}

	return nil
}

func (fileMock *FileMock) SyscallConn() (syscall.RawConn, error) {
	fileMock.Mock.SyscallConn.AddInput(struct{}{})

	result, err := fileMock.Mock.SyscallConn.GetNextResult()
	if err != nil {
		return nil, err
	}

	return *result, nil
}

func (fileMock *FileMock) Truncate(size int64) error {
	fileMock.Mock.Truncate.AddInput(struct {
		Size int64
	}{
		Size: size,
	})

	result, err := fileMock.Mock.Truncate.GetNextResult()
	if err != nil {
		return err
	}

	if *result {
		return fmt.Errorf("Error requested")
	}

	return nil
}

func (fileMock *FileMock) Write(b []byte) (n int, err error) {
	fileMock.Mock.Write.AddInput(struct {
		B []byte
	}{
		B: b,
	})

	result, err := fileMock.Mock.Write.GetNextResult()
	if err != nil {
		return 0, err
	}

	return *result, nil
}

func (fileMock *FileMock) WriteAt(b []byte, off int64) (n int, err error) {
	fileMock.Mock.WriteAt.AddInput(struct {
		B   []byte
		Off int64
	}{
		B:   b,
		Off: off,
	})

	result, err := fileMock.Mock.WriteAt.GetNextResult()
	if err != nil {
		return 0, err
	}

	return *result, nil
}

func (fileMock *FileMock) WriteString(s string) (n int, err error) {
	fileMock.Mock.WriteString.AddInput(struct {
		S string
	}{
		S: s,
	})

	result, err := fileMock.Mock.WriteString.GetNextResult()
	if err != nil {
		return 0, err
	}

	return *result, nil
}

func (fileMock *FileMock) WriteTo(w io.Writer) (n int64, err error) {
	fileMock.Mock.WriteTo.AddInput(struct {
		W io.Writer
	}{
		W: w,
	})

	result, err := fileMock.Mock.WriteTo.GetNextResult()
	if err != nil {
		return 0, err
	}

	return *result, nil
}
