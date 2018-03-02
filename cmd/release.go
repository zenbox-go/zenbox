package cmd

import "gopkg.in/urfave/cli.v2"

var CmdRelease = &cli.Command{
	Name:  "release",
	Usage: "构建 Go 生产环境版本",
}
