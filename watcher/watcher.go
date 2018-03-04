package watcher

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"zenbox/config"
	"zenbox/print"

	"github.com/fsnotify/fsnotify"
)

func Sync(c context.Context) error {
	cfg, ok := c.Value("CONFIG").(*config.Config)
	if !ok {
		return errors.New("配置文件错误")
	}

	ctx, cancel := context.WithCancel(c)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("创建热编译服务错误: %v", err)
	}
	defer watcher.Close()

	for _, dir := range cfg.Watcher.Dirs {
		_ = watcher.Add(dir)
	}

	print.I("热编译服务正在运行...")

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case err := <-watcher.Errors:
				print.EF("文件监控错误: %v", err)
			case event := <-watcher.Events:
				// 忽略修改权限操作,因为这不影响编译
				if event.Op&fsnotify.Chmod == fsnotify.Chmod {
					continue
				}

				// 忽略测试文件,因为测试文件不需要重新编译
				if checkTestFile(event.Name) {
					continue
				}

				ext := filepath.Ext(event.Name)
				if !checkExt(ext, cfg.Watcher.Exts) {
					continue
				}

				// 当删除目录时,需要将目录从监控目录列表删除
				if event.Op&fsnotify.Remove == fsnotify.Remove {
					if fi, err := os.Stat(event.Name); err == nil && fi.IsDir() {
						_ = watcher.Remove(event.Name)
					}
				}

				fmt.Printf("%v: %s\n", event.Op, event.Name)
				fmt.Println("================================")
				fmt.Println("开始编译...")

				go build(ctx)
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	cancel()

	print.W("热编译服务结束运行.")

	return nil
}

func build(c context.Context) {

}

func checkTestFile(name string) bool {
	return strings.HasSuffix(filepath.Base(name), "_test.go")
}

func checkExt(name string, exts []string) bool {
	for _, ext := range exts {
		if ext == name {
			return true
		}
	}

	return false
}
