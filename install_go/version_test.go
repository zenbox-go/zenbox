package install_go

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"testing"
)

func TestGetGoVersionsFormRemote(t *testing.T) {
	err := getGoVersionsFormRemote()
	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadFile("cache/VERSION")
	if err != nil {
		t.Fatal(err)
	}

	vs := make([]string, 0)

	scanner := bufio.NewScanner(bytes.NewReader(b))
	for scanner.Scan() {
		vs = append(vs, scanner.Text())
	}

	t.Log(vs)
}

func TestGetGoVersions(t *testing.T) {
	vs, err := getGoVersions()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(vs)
}
