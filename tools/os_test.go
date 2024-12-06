package tools_test

import (
	"testing"

	"github.com/nextunit-io/go-tools/tools"
	"github.com/nextunit-io/go-tools/toolsmock"
	"github.com/stretchr/testify/assert"
)

func TestOsInstanceReturnsNotNil(t *testing.T) {
	tools.SetOsInstance(nil)
	assert.NotNil(t, tools.GetOsInstance())
}

func TestOwnOsClient(t *testing.T) {
	client := toolsmock.GetOsMock()
	tools.SetOsInstance(client)
	assert.Equal(t, client, tools.GetOsInstance())
}
