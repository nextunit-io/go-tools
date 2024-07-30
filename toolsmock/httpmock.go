package toolsmock

import (
	"fmt"
	"net/http"

	gomock "github.com/nextunit-io/go-mock"
)

type httpMockStruct struct {
	Get *gomock.ToolMock[struct {
		Url string
	},
		http.Response,
	]
	Delete *gomock.ToolMock[struct {
		Url string
	},
		http.Response,
	]
	Post *gomock.ToolMock[struct {
		Url string
	},
		http.Response,
	]
	Put *gomock.ToolMock[struct {
		Url string
	},
		http.Response,
	]
}

type httpMock struct {
	Mock httpMockStruct
}

func GetHttpMock() *httpMock {
	return &httpMock{
		Mock: httpMockStruct{
			Get: gomock.GetMock[struct {
				Url string
			},
				http.Response,
			](fmt.Errorf("GET general error")),
			Put: gomock.GetMock[struct {
				Url string
			},
				http.Response,
			](fmt.Errorf("PUT general error")),
			Post: gomock.GetMock[struct {
				Url string
			},
				http.Response,
			](fmt.Errorf("POST general error")),
			Delete: gomock.GetMock[struct {
				Url string
			},
				http.Response,
			](fmt.Errorf("DELETE general error")),
		},
	}
}

func (h *httpMock) Get(url string) (resp *http.Response, err error) {
	h.Mock.Get.AddInput(
		struct {
			Url string
		}{
			Url: url,
		},
	)

	return h.Mock.Get.GetNextResult()
}

func (h *httpMock) Put(url string) (resp *http.Response, err error) {
	h.Mock.Put.AddInput(
		struct {
			Url string
		}{
			Url: url,
		},
	)

	return h.Mock.Put.GetNextResult()
}

func (h *httpMock) Post(url string) (resp *http.Response, err error) {
	h.Mock.Post.AddInput(
		struct {
			Url string
		}{
			Url: url,
		},
	)

	return h.Mock.Post.GetNextResult()
}

func (h *httpMock) Delete(url string) (resp *http.Response, err error) {
	h.Mock.Delete.AddInput(
		struct {
			Url string
		}{
			Url: url,
		},
	)

	return h.Mock.Delete.GetNextResult()
}
