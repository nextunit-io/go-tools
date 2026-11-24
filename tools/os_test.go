package tools_test

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
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

func TestDefaultOsClientCreateTemp(t *testing.T) {
	tools.SetOsInstance(nil)
	dir := t.TempDir()

	file, err := tools.GetOsInstance().CreateTemp(dir, "prefix-*.txt")
	assert.Nil(t, err)
	assert.NotNil(t, file)
	defer file.Close()

	name := file.Name()
	assert.Equal(t, dir, filepath.Dir(name))
	base := filepath.Base(name)
	assert.True(t, strings.HasPrefix(base, "prefix-"))
	assert.True(t, strings.HasSuffix(base, ".txt"))

	info, err := os.Stat(name)
	assert.Nil(t, err)
	assert.Equal(t, os.FileMode(0600), info.Mode().Perm())
}

func TestDefaultOsClientCreateTempMissingDir(t *testing.T) {
	tools.SetOsInstance(nil)
	dir := filepath.Join(t.TempDir(), "missing")

	file, err := tools.GetOsInstance().CreateTemp(dir, "prefix-*")
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, fs.ErrNotExist)
	assert.Nil(t, file)
}

func TestDefaultOsClientLink(t *testing.T) {
	tools.SetOsInstance(nil)
	dir := t.TempDir()
	oldname := filepath.Join(dir, "old")
	newname := filepath.Join(dir, "new")
	assert.Nil(t, os.WriteFile(oldname, []byte("content"), 0644))

	assert.Nil(t, tools.GetOsInstance().Link(oldname, newname))

	content, err := os.ReadFile(newname)
	assert.Nil(t, err)
	assert.Equal(t, []byte("content"), content)

	oldInfo, err := os.Stat(oldname)
	assert.Nil(t, err)
	newInfo, err := os.Stat(newname)
	assert.Nil(t, err)
	assert.True(t, os.SameFile(oldInfo, newInfo))
}

func TestDefaultOsClientLinkExistingTarget(t *testing.T) {
	tools.SetOsInstance(nil)
	dir := t.TempDir()
	oldname := filepath.Join(dir, "old")
	newname := filepath.Join(dir, "new")
	assert.Nil(t, os.WriteFile(oldname, []byte("old"), 0644))
	assert.Nil(t, os.WriteFile(newname, []byte("new"), 0644))

	err := tools.GetOsInstance().Link(oldname, newname)
	assert.ErrorIs(t, err, fs.ErrExist)
}

func TestDefaultOsClientLinkMissingSource(t *testing.T) {
	tools.SetOsInstance(nil)
	dir := t.TempDir()

	err := tools.GetOsInstance().Link(filepath.Join(dir, "missing"), filepath.Join(dir, "new"))
	assert.ErrorIs(t, err, fs.ErrNotExist)
}

func TestDefaultOsClientRename(t *testing.T) {
	tools.SetOsInstance(nil)
	dir := t.TempDir()
	oldpath := filepath.Join(dir, "old")
	newpath := filepath.Join(dir, "new")
	assert.Nil(t, os.WriteFile(oldpath, []byte("content"), 0644))

	assert.Nil(t, tools.GetOsInstance().Rename(oldpath, newpath))

	_, err := os.Stat(oldpath)
	assert.ErrorIs(t, err, fs.ErrNotExist)
	content, err := os.ReadFile(newpath)
	assert.Nil(t, err)
	assert.Equal(t, []byte("content"), content)
}

func TestDefaultOsClientRenameMissingSource(t *testing.T) {
	tools.SetOsInstance(nil)
	dir := t.TempDir()

	err := tools.GetOsInstance().Rename(filepath.Join(dir, "missing"), filepath.Join(dir, "new"))
	assert.ErrorIs(t, err, fs.ErrNotExist)
}
