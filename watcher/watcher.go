package watcher

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)

func Sync(c context.Context) error {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("创建热编译服务错误: %v", err)
	}
	defer watcher.Close()

	if err = filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			_ = watcher.Add(path)
		}

		return nil
	}); err != nil {

		return fmt.Errorf("遍历当前目录错误: %v", err)
	}

	go func() {
		for err := range watcher.Errors {
			fmt.Println(err)
		}
	}()

	fmt.Println("热编译服务正在运行...")

	for event := range watcher.Events {
		ext := filepath.Ext(event.Name)
		if ext == ".tmp" || ext == ".temp" || ext == ".swp" || ext == ext+"~" {
			continue
		}

		if strings.HasSuffix(event.Name, "___jb_tmp___") ||
			strings.HasSuffix(event.Name, "___jb_old___") {
			continue
		}

		fi, err := os.Stat(event.Name)
		if err == nil && fi.IsDir() {
			continue
		}

		fmt.Println(event)

		switch {
		case event.Op&fsnotify.Create == fsnotify.Create:
			fmt.Printf("%v %v\n", event.Op, event.Name)
		case event.Op&fsnotify.Write == fsnotify.Write,
			event.Op&fsnotify.Chmod == fsnotify.Chmod:
			fmt.Printf("%v %v\n", event.Op, event.Name)
		case event.Op&fsnotify.Remove == fsnotify.Remove:
			fmt.Printf("%v %v\n", event.Op, event.Name)
		case event.Op&fsnotify.Rename == fsnotify.Rename:
			fmt.Printf("%v %v\n", event.Op, event.Name)
		default:
			panic(fmt.Sprintf("未知的文件操作类型: %v", event.Op))
		}
	}

	fmt.Println("热编译服务结束运行.")

	return nil
}
