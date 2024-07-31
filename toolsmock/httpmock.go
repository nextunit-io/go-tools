package toolsmock

import (
	"fmt"
	"io"
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
		Url         string
		ContentType string
		Body        io.Reader
	},
		http.Response,
	]
	Put *gomock.ToolMock[struct {
		Url         string
		ContentType string
		Body        io.Reader
	},
		http.Response,
	]
}

type HttpMock struct {
	Mock httpMockStruct
}

func GetHttpMock() *HttpMock {
	return &HttpMock{
		Mock: httpMockStruct{
			Get: gomock.GetMock[struct {
				Url string
			},
				http.Response,
			](fmt.Errorf("GET general error")),
			Put: gomock.GetMock[struct {
				Url         string
				ContentType string
				Body        io.Reader
			},
				http.Response,
			](fmt.Errorf("PUT general error")),
			Post: gomock.GetMock[struct {
				Url         string
				ContentType string
				Body        io.Reader
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

func (h *HttpMock) Get(url string) (resp *http.Response, err error) {
	h.Mock.Get.AddInput(
		struct {
			Url string
		}{
			Url: url,
		},
	)

	return h.Mock.Get.GetNextResult()
}

func (h *HttpMock) Put(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	h.Mock.Put.AddInput(
		struct {
			Url         string
			ContentType string
			Body        io.Reader
		}{
			Url:         url,
			ContentType: contentType,
			Body:        body,
		},
	)

	return h.Mock.Put.GetNextResult()
}

func (h *HttpMock) Post(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	h.Mock.Post.AddInput(
		struct {
			Url         string
			ContentType string
			Body        io.Reader
		}{
			Url:         url,
			ContentType: contentType,
			Body:        body,
		},
	)

	return h.Mock.Post.GetNextResult()
}

func (h *HttpMock) Delete(url string) (resp *http.Response, err error) {
	h.Mock.Delete.AddInput(
		struct {
			Url string
		}{
			Url: url,
		},
	)

	return h.Mock.Delete.GetNextResult()
}
