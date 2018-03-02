package watcher

import (
	"path/filepath"
	"testing"
)

func TestTemp(t *testing.T) {
	t.Log(filepath.Ext("aaa.tmp"))
}
