package main

import (
	"fmt"
	"os"
	"time"

	"zenbox/cmd"

	"gopkg.in/urfave/cli.v2"
)

var (
	Version = "0.1"
)

func main() {
	app := &cli.App{
		Name:     "zenbox",
		Usage:    "做好用的 Go 项目管理工具",
		Version:  Version,
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "果子",
				Email: "zmguozi@gmail.com",
			},
		},
		HelpName:              "zenbox",
		HideHelp:              true,
		EnableShellCompletion: true,
		ArgsUsage:             "[arrgh]",
		Commands: []*cli.Command{
			cmd.CmdInit,
			cmd.CmdTools,
			cmd.CmdSearch,
			cmd.CmdWatch,
			cmd.CmdRelease,
			cmd.CmdSetup,
			cmd.CmdClear,
			cmd.CmdHelp,
		},
		CommandNotFound: func(c *cli.Context, command string) {
			fmt.Fprintf(c.App.Writer, "没有这个指令 '%s'\n", command)
		},
		OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
			if isSubcommand {
				return err
			}

			fmt.Fprintf(c.App.Writer, "没有这个参数: %v\n", c.NumFlags())
			return nil
		},
	}

	app.Flags = append(app.Flags, []cli.Flag{}...)
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
