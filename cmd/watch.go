package cmd

import (
	"context"
	"zenbox/watcher"

	"gopkg.in/urfave/cli.v2"
)

var CmdWatch = &cli.Command{
	Name:   "watch",
	Usage:  "监控 Go 源码变化并自动编译和运行",
	Action: watchAction,
}

func watchAction(c *cli.Context) error {
	return watcher.Sync(context.Background())
}
