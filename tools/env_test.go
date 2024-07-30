package tools_test

import (
	"testing"

	"github.com/nextunit-io/go-tools/tools"
	"github.com/nextunit-io/go-tools/toolsmock"
	"github.com/stretchr/testify/assert"
)

func TestGetInstanceReturnsNotNil(t *testing.T) {
	tools.SetEnvGetInstance(nil)
	assert.NotNil(t, tools.GetEnvInstance())
}

func TestOwnEnvClient(t *testing.T) {
	client := toolsmock.GetEnvMock()
	tools.SetEnvGetInstance(client)
	assert.Equal(t, client, tools.GetEnvInstance())
}
