package tools_test

import (
	"testing"

	"github.com/nextunit-io/go-tools/tools"
	"github.com/nextunit-io/go-tools/toolsmock"
	"github.com/stretchr/testify/assert"
)

func TestGetHttpInstanceReturnsNotNil(t *testing.T) {
	tools.SetHttpCreator(nil)
	auth := "test-auth"
	assert.NotNil(t, tools.GetHttpInstance(tools.HttpConfig{
		Authorization: &auth,
	}))
}

func TestOwnHttpClient(t *testing.T) {
	client := toolsmock.GetHttpMock()
	tools.SetHttpCreator(func(cfg tools.HttpConfig) tools.Http {
		return client
	})
	auth := "test-auth"
	assert.Equal(t, client, tools.GetHttpInstance(tools.HttpConfig{
		Authorization: &auth,
	}))
}
