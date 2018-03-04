package install_go

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"
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

	if err := downloadGolang(
		"go1.8.1",
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
	m, err := getGoVersions()
	if err != nil {
		t.Fatal(err)
	}

	for _, value := range m {
		t.Log(value)
	}
}

func TestPingURL(t *testing.T) {
	// 超时3秒即代表地址无法访问
	timeout := time.Duration(3 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Head("https://www.google.com")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	t.Log("可以连接")
}
