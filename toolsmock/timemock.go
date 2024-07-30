package toolsmock

import (
	"fmt"
	"time"

	gomock "github.com/nextunit-io/go-mock"
)

type timeMockStruct struct {
	Now *gomock.ToolMock[
		interface{},
		time.Time,
	]
}

type timeMock struct {
	Mock timeMockStruct
}

func GetTimeMock() *timeMock {
	return &timeMock{
		Mock: timeMockStruct{
			Now: gomock.GetMock[
				interface{},
				time.Time,
			](fmt.Errorf("NOW general error")),
		},
	}
}

func (t *timeMock) Now() time.Time {
	t.Mock.Now.AddInput(nil)

	result, err := t.Mock.Now.GetNextResult()
	if err != nil {
		panic(err.Error())
	}

	return *result
}
