package cmd

import (
	"context"
	"fmt"
	"zenbox/config"
	"zenbox/watcher"

	"gopkg.in/urfave/cli.v2"
)

var CmdWatch = &cli.Command{
	Name:   "watch",
	Usage:  "监控 Go 源码变化并自动编译和运行",
	Action: watchAction,
}

func watchAction(c *cli.Context) error {
	cfg, err := config.New()
	if err != nil {
		return fmt.Errorf("初始化监控配置文件错误: %v\n", err)
	}

	ctx := context.WithValue(context.Background(), "CONFIG", cfg)
	return watcher.Sync(ctx)
}
