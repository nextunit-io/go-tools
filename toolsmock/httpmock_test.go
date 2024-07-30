package toolsmock_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/nextunit-io/go-tools/toolsmock"
	"github.com/stretchr/testify/assert"
)

func TestHttpMockGet(t *testing.T) {
	t.Helper()
	httpMock := toolsmock.GetHttpMock()

	t.Run("Testing Get", func(t *testing.T) {
		output := http.Response{
			StatusCode: 200,
		}

		httpMock.Mock.Get.AddReturnValue(&output)
		httpMock.Mock.Get.AddReturnValue(&output)
		httpMock.Mock.Get.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			o, err := httpMock.Get(fmt.Sprintf("test-input-%d", i))

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := httpMock.Get("test-input-error")

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("GET general error"), err)

		assert.Equal(t, 4, httpMock.Mock.Get.HasBeenCalled())
		for i := 0; i < 3; i++ {
			input := httpMock.Mock.Get.GetInput(i)
			assert.Equal(t, fmt.Sprintf("test-input-%d", i), input.Url)
		}

		input := httpMock.Mock.Get.GetInput(3)
		assert.Equal(t, "test-input-error", input.Url)
	})
}

func TestHttpMockPut(t *testing.T) {
	t.Helper()
	httpMock := toolsmock.GetHttpMock()

	t.Run("Testing Put", func(t *testing.T) {
		output := http.Response{
			StatusCode: 200,
		}

		httpMock.Mock.Put.AddReturnValue(&output)
		httpMock.Mock.Put.AddReturnValue(&output)
		httpMock.Mock.Put.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			o, err := httpMock.Put(fmt.Sprintf("test-input-%d", i))

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := httpMock.Put("test-input-error")

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("PUT general error"), err)

		assert.Equal(t, 4, httpMock.Mock.Put.HasBeenCalled())
		for i := 0; i < 3; i++ {
			input := httpMock.Mock.Put.GetInput(i)
			assert.Equal(t, fmt.Sprintf("test-input-%d", i), input.Url)
		}

		input := httpMock.Mock.Put.GetInput(3)
		assert.Equal(t, "test-input-error", input.Url)
	})
}
func TestHttpMockPost(t *testing.T) {
	t.Helper()
	httpMock := toolsmock.GetHttpMock()

	t.Run("Testing Post", func(t *testing.T) {
		output := http.Response{
			StatusCode: 200,
		}

		httpMock.Mock.Post.AddReturnValue(&output)
		httpMock.Mock.Post.AddReturnValue(&output)
		httpMock.Mock.Post.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			o, err := httpMock.Post(fmt.Sprintf("test-input-%d", i))

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := httpMock.Post("test-input-error")

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("POST general error"), err)

		assert.Equal(t, 4, httpMock.Mock.Post.HasBeenCalled())
		for i := 0; i < 3; i++ {
			input := httpMock.Mock.Post.GetInput(i)
			assert.Equal(t, fmt.Sprintf("test-input-%d", i), input.Url)
		}

		input := httpMock.Mock.Post.GetInput(3)
		assert.Equal(t, "test-input-error", input.Url)
	})
}

func TestHttpMockDelete(t *testing.T) {
	t.Helper()
	httpMock := toolsmock.GetHttpMock()

	t.Run("Testing Delete", func(t *testing.T) {
		output := http.Response{
			StatusCode: 200,
		}

		httpMock.Mock.Delete.AddReturnValue(&output)
		httpMock.Mock.Delete.AddReturnValue(&output)
		httpMock.Mock.Delete.AddReturnValue(&output)

		for i := 0; i < 3; i++ {
			o, err := httpMock.Delete(fmt.Sprintf("test-input-%d", i))

			assert.Nil(t, err)
			assert.Equal(t, output, *o)
		}

		o, err := httpMock.Delete("test-input-error")

		assert.Nil(t, o)
		assert.Equal(t, fmt.Errorf("DELETE general error"), err)

		assert.Equal(t, 4, httpMock.Mock.Delete.HasBeenCalled())
		for i := 0; i < 3; i++ {
			input := httpMock.Mock.Delete.GetInput(i)
			assert.Equal(t, fmt.Sprintf("test-input-%d", i), input.Url)
		}

		input := httpMock.Mock.Delete.GetInput(3)
		assert.Equal(t, "test-input-error", input.Url)
	})
}
