package config

import "testing"

func TestNew(t *testing.T) {
	cfg, err := New()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(cfg.Watcher)
}
