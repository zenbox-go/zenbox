package cmd

import (
	"context"
	"zenbox/install_go"

	"gopkg.in/urfave/cli.v2"
)

var CmdSetup = &cli.Command{
	Name:   "setup",
	Usage:  "安装 Go 开发环境,切换 Go 版本",
	Action: setupAction,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "proxy",
			Usage: "设置代理地址",
		},
		&cli.StringFlag{
			Name:  "download_url",
			Usage: "Go 安装包下载地址前缀",
		},
		&cli.StringFlag{
			Name:  "version_url",
			Usage: "Go 版本号列表地址",
		},
	},
}

func setupAction(c *cli.Context) error {

	if c.IsSet("proxy") {
		install_go.DefaultProxyURL = c.String("proxy")
	}
	if c.IsSet("download_url") {
		install_go.DefaultDownloadURLPrefix = c.String("download_url")
	}
	if c.IsSet("version_url") {
		install_go.DefaultSourceURL = c.String("version_url")
	}

	install_go.Setup(context.Background())

	return nil
}
