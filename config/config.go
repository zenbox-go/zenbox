package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Watcher watcher `toml:"watcher"`
}

type watcher struct {
	Dirs    []string `toml:"dirs"`
	Exts    []string `toml:"exts"`
	Ignores []string `toml:"ignores"`
	a       byte
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
	if len(watcher.Exts) == 0 {
		watcher.Exts = append(watcher.Exts, ".go", ".gohtml")
	}

	if len(watcher.Ignores) == 0 {
		watcher.Ignores = append(watcher.Ignores, "vendor", ".idea")
	}

	if len(watcher.Dirs) == 0 {
		if err = filepath.Walk(currentPath, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				if len(watcher.Ignores) > 0 {
					for _, ignore := range watcher.Ignores {
						if ignore != path {
							watcher.Dirs = append(watcher.Dirs, path)
						}
					}
				} else {
					watcher.Dirs = append(watcher.Dirs, path)
				}
			}

			return nil
		}); err != nil {
			return nil, fmt.Errorf("遍历当前目录错误: %v", err)
		}
	}

	return cfg, nil
}
