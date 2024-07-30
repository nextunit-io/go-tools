package toolsmock

import (
	"fmt"

	gomock "github.com/nextunit-io/go-mock"
)

type envMockStruct struct {
	Getenv *gomock.ToolMock[struct {
		Key string
	},
		string,
	]
}

type envMock struct {
	Mock envMockStruct
}

func GetEnvMock() *envMock {
	return &envMock{
		Mock: envMockStruct{
			Getenv: gomock.GetMock[struct {
				Key string
			},
				string,
			](fmt.Errorf("PutEvents general error")),
		},
	}
}

func (e *envMock) Getenv(key string) string {
	e.Mock.Getenv.AddInput(
		struct {
			Key string
		}{
			Key: key,
		},
	)

	result, err := e.Mock.Getenv.GetNextResult()
	if err != nil {
		panic(err.Error())
	}

	return *result
}
