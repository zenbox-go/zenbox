package install_go

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/manifoldco/promptui"
)

type step func(context.Context) error

func welcome(_ context.Context) error {
	fmt.Println("欢迎使用 Go 安装向导!")

	p1 := promptui.Prompt{
		Label:     "是否需要安装 Go ",
		IsConfirm: true,
	}

	result, err := p1.Run()
	if err != nil {
		return exitSetup
	}

	if strings.ToLower(result) != "y" {
		return exitSetup
	}

	return nil
}

func checkGoInstalled(ctx context.Context) error {
	path, err := whichGo(ctx)
	if err != nil {
		fmt.Printf("检查 Go 是否应安装错误: %v\n", err)
	}

	if path == "" {
		return nil
	}

	if path != DefaultInstallPath {
		fmt.Printf("\n检测到 %s 已安装,安装路径: [%v]\n\n", getLocalGoVersion(ctx), getGOROOT(ctx))
	}

	return nil
}

func setupGo(_ context.Context) error {
	fmt.Print("正在获取Golang版本号列表...")
	versions, err := getGoVersions()
	if err != nil {
		return err
	}

	templates := &promptui.SelectTemplates{
		Help: `{{ "按[CTRL+C]退出安装,使用方向键选择安装版本:" | faint }} {{ .NextKey | faint }} ` +
			`{{ .PrevKey | faint }} {{ .PageDownKey | faint }} {{ .PageUpKey | faint }} ` +
			`{{ if .Search }} {{ "and" | faint }} {{ .SearchKey | faint }} {{ "toggles search" | faint }}{{ end }}`,
	}

	prompt := &promptui.Select{
		Label:     "选择需要安装的 Go 版本",
		Items:     versions,
		Size:      10,
		Templates: templates,
	}

	_, version, err := prompt.Run()
	if err != nil {
		return exitSetup
	}

	suffix := "tar.gz"
	if runtime.GOOS == "windows" {
		suffix = "zip"
	}

	targetFile := fmt.Sprintf("go%s.%s-%s.%s", version, runtime.GOOS, runtime.GOARCH, suffix)
	cacheFile := filepath.Join("cache", "downloads", targetFile)

	if _, e := os.Stat(cacheFile); os.IsNotExist(e) {
		if err := downloadGolang(targetFile, DefaultInstallPath); err != nil {
			return err
		}
	} else {
		unpackFn := unpackTar
		if runtime.GOOS == "windows" {
			unpackFn = unpackZip
		}

		os.RemoveAll(DefaultInstallPath)

		fmt.Printf("正在安装 Go 到路径: [%s]\n", DefaultInstallPath)
		if err := unpackFn(cacheFile, DefaultInstallPath); err != nil {
			return fmt.Errorf("解压 Go 到目标路径 %s 失败: %v", DefaultInstallPath, err)
		}
	}

	// set $GOROOT
	if err := persistEnvVar("GOROOT", DefaultInstallPath); err != nil {
		return err
	}

	if err := appendToPATH(filepath.Join(DefaultInstallPath, "bin")); err != nil {
		return err
	}

	fmt.Printf("\nGo%s 安装完成,安装路径: [%s]\n\n", version, DefaultInstallPath)

	return nil
}

func setGOPATH(_ context.Context) error {
	home, err := getHomeDir()
	if err != nil {
		return err
	}

	gopath := filepath.Join(home, "go")

	p2 := promptui.Prompt{
		Label:   "输入工作目录路径[GOPATH]",
		Default: gopath,
	}

	gopath, err = p2.Run()
	if err != nil {
		return err
	}

	if err := persistEnvVar("GOPATH", gopath); err != nil {
		return err
	}

	if err := appendToPATH(filepath.Join(gopath, "bin")); err != nil {
		return err
	}

	fmt.Printf("工作目录路径[GOPATH]: %s 设置完成!\n", gopath)

	return nil
}

func setupDone(_ context.Context) error {
	return persistEnvChangesForSession()
}
