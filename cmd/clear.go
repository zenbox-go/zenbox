package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"gopkg.in/urfave/cli.v2"
)

var CmdClear = &cli.Command{
	Name:   "clear",
	Usage:  "清理缓存",
	Action: clearAction,
}

func clearAction(_ *cli.Context) error {
	prompt := promptui.Prompt{
		Label:     "是否清理缓存",
		IsConfirm: true,
	}

	y, err := prompt.Run()
	if err != nil {
		return errors.New("放弃清理缓存")
	}

	if strings.ToLower(y) != "y" {
		return errors.New("放弃清理缓存")
	}

	if err = os.RemoveAll("cache"); err != nil {
		return cli.Exit("缓存清理失败: "+err.Error(), 1)
	}

	fmt.Fprintln(os.Stdout, "缓存清理完毕")

	return nil
}
