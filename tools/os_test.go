package tools_test

import (
	"fmt"
	"testing"

	"github.com/nextunit-io/go-tools/interfaces"
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

func TestOsClientCreateTemp(t *testing.T) {
	client := toolsmock.GetOsMock()
	tools.SetOsInstance(client)

	var file interfaces.OsFileInterface = toolsmock.GetFileMock()
	client.Mock.CreateTemp.AddReturnValue(&file)

	result, err := tools.GetOsInstance().CreateTemp("test-dir", "test-*.tmp")
	assert.Nil(t, err)
	assert.Equal(t, file, result)

	result, err = tools.GetOsInstance().CreateTemp("test-dir-error", "test-error-*.tmp")
	assert.Nil(t, result)
	assert.Equal(t, fmt.Errorf("CreateTemp general error"), err)

	assert.Equal(t, 2, client.Mock.CreateTemp.HasBeenCalled())
	assert.Equal(t, "test-dir", client.Mock.CreateTemp.GetInput(0).Dir)
	assert.Equal(t, "test-*.tmp", client.Mock.CreateTemp.GetInput(0).Pattern)
	assert.Equal(t, "test-dir-error", client.Mock.CreateTemp.GetInput(1).Dir)
	assert.Equal(t, "test-error-*.tmp", client.Mock.CreateTemp.GetInput(1).Pattern)
}

func TestOsClientLink(t *testing.T) {
	client := toolsmock.GetOsMock()
	tools.SetOsInstance(client)

	noError := false
	errorRequested := true
	client.Mock.Link.AddReturnValue(&noError)
	client.Mock.Link.AddReturnValue(&errorRequested)

	assert.Nil(t, tools.GetOsInstance().Link("test-old", "test-new"))
	assert.Equal(t, fmt.Errorf("Error requested"), tools.GetOsInstance().Link("test-old-requested", "test-new-requested"))
	assert.Equal(t, fmt.Errorf("Link general error"), tools.GetOsInstance().Link("test-old-error", "test-new-error"))

	assert.Equal(t, 3, client.Mock.Link.HasBeenCalled())
	assert.Equal(t, "test-old", client.Mock.Link.GetInput(0).Oldname)
	assert.Equal(t, "test-new", client.Mock.Link.GetInput(0).Newname)
	assert.Equal(t, "test-old-requested", client.Mock.Link.GetInput(1).Oldname)
	assert.Equal(t, "test-new-requested", client.Mock.Link.GetInput(1).Newname)
	assert.Equal(t, "test-old-error", client.Mock.Link.GetInput(2).Oldname)
	assert.Equal(t, "test-new-error", client.Mock.Link.GetInput(2).Newname)
}

func TestOsClientRename(t *testing.T) {
	client := toolsmock.GetOsMock()
	tools.SetOsInstance(client)

	noError := false
	errorRequested := true
	client.Mock.Rename.AddReturnValue(&noError)
	client.Mock.Rename.AddReturnValue(&errorRequested)

	assert.Nil(t, tools.GetOsInstance().Rename("test-old", "test-new"))
	assert.Equal(t, fmt.Errorf("Error requested"), tools.GetOsInstance().Rename("test-old-requested", "test-new-requested"))
	assert.Equal(t, fmt.Errorf("Rename general error"), tools.GetOsInstance().Rename("test-old-error", "test-new-error"))

	assert.Equal(t, 3, client.Mock.Rename.HasBeenCalled())
	assert.Equal(t, "test-old", client.Mock.Rename.GetInput(0).Oldpath)
	assert.Equal(t, "test-new", client.Mock.Rename.GetInput(0).Newpath)
	assert.Equal(t, "test-old-requested", client.Mock.Rename.GetInput(1).Oldpath)
	assert.Equal(t, "test-new-requested", client.Mock.Rename.GetInput(1).Newpath)
	assert.Equal(t, "test-old-error", client.Mock.Rename.GetInput(2).Oldpath)
	assert.Equal(t, "test-new-error", client.Mock.Rename.GetInput(2).Newpath)
}
