// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris

package install_go

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
)

const (
	envSeparator = ":"
	homeKey      = "HOME"
	lineEnding   = "\n"
	pathVar      = "$PATH"
)

var DefaultInstallPath = func() string {
	home, err := getHomeDir()
	if err != nil {
		return "/usr/local/go"
	}

	return filepath.Join(home, ".go")
}()

func whichGo(ctx context.Context) (string, error) {
	return findGo(ctx, "which")
}

func currentShell() string {
	return os.Getenv("SHELL")
}

func persistEnvChangesForSession() error {
	shellConfig, err := shellConfigFile()
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Printf("现在请执行 `source %s` 命令,将环境变量作用于当前终端会话.\n", shellConfig)
	fmt.Println("或者关闭此终端,重新启动一个终端来执行 Go 命令.")

	return nil
}
