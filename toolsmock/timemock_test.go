package toolsmock_test

import (
	"testing"
	"time"

	"github.com/nextunit-io/go-tools/toolsmock"
	"github.com/stretchr/testify/assert"
)

func TestTimeMockNow(t *testing.T) {
	t.Helper()
	timeMock := toolsmock.GetTimeMock()

	t.Run("Testing Now", func(t *testing.T) {
		checkTime := time.Date(2024, 12, 01, 00, 00, 00, 0, time.UTC)

		timeMock.Mock.Now.AddReturnValue(&checkTime)
		timeMock.Mock.Now.AddReturnValue(&checkTime)
		timeMock.Mock.Now.AddReturnValue(&checkTime)

		for i := 0; i < 3; i++ {
			o := timeMock.Now()
			assert.Equal(t, checkTime, o)
		}

		assert.Equal(t, 3, timeMock.Mock.Now.HasBeenCalled())
	})
}
