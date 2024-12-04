package tools_test

import (
	"testing"

	"github.com/nextunit-io/go-tools/tools"
	"github.com/nextunit-io/go-tools/toolsmock"
	"github.com/stretchr/testify/assert"
)

func TestExecInstanceReturnsNotNil(t *testing.T) {
	tools.SetExecInstance(nil)
	assert.NotNil(t, tools.GetExecInstance())
}

func TestOwnExecClient(t *testing.T) {
	client := toolsmock.GetExecMock()
	tools.SetExecInstance(client)
	assert.Equal(t, client, tools.GetExecInstance())
}
