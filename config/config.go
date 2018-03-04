package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Watcher *watcher `toml:"watcher"`
}

type watcher struct {
	Dirs    []string `toml:"dirs"`
	Exts    []string `toml:"exts"`
	Ignores []string `toml:"ignores"`
}

func New(name ...string) (*Config, error) {
	currentPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(currentPath, ".zenbox.toml")
	if len(name) > 0 {
		configPath = name[0]
	}

	var cfg *Config
	if _, err := toml.DecodeFile(configPath, &cfg); err != nil {
		return nil, err
	}

	watcher := cfg.Watcher
	if err = filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			return nil
		}

		if strings.HasPrefix(path, ".git") ||
			strings.HasPrefix(path, "vendor") ||
			strings.HasPrefix(path, ".idea") {
			return nil
		}

		for _, dir := range watcher.Ignores {
			if path == dir {
				return nil
			}
		}

		watcher.Dirs = append(watcher.Dirs, path)

		return nil
	}); err != nil {
		return nil, fmt.Errorf("遍历当前目录错误: %v", err)
	}

	return cfg, nil
}
