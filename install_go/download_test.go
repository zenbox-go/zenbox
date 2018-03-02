package install_go

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestDownloadGoVersion(t *testing.T) {
	if testing.Short() {
		t.Skipf("Skipping download in short mode")
	}

	tmpd, err := ioutil.TempDir("", "go")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpd)

	if err := downloadGoVersion(
		"go1.8.1",
		"linux",
		"amd64",
		filepath.Join(tmpd, "go"),
	); err != nil {
		t.Fatal(err)
	}

	vf := filepath.Join(tmpd, "go", "VERSION")
	if _, err := os.Stat(vf); os.IsNotExist(err) {
		t.Fatalf("file %s does not exist and should", vf)
	}
}

func TestGetAllGoVersion(t *testing.T) {
	m, err := getAllGoVersion()
	if err != nil {
		t.Fatal(err)
	}

	for _, value := range m {
		t.Log(value)
	}
}
