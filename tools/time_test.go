package tools_test

import (
	"testing"

	"github.com/nextunit-io/go-tools/tools"
	"github.com/nextunit-io/go-tools/toolsmock"
	"github.com/stretchr/testify/assert"
)

func TestTimeGetInstanceReturnsNotNil(t *testing.T) {
	tools.SetTimeInstance(nil)
	assert.NotNil(t, tools.GetTimeInstance())
}

func TestOwnTimeClient(t *testing.T) {
	client := toolsmock.GetTimeMock()
	tools.SetTimeInstance(client)
	assert.Equal(t, client, tools.GetTimeInstance())
}
