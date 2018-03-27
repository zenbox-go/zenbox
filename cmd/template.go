package cmd

import "gopkg.in/urfave/cli.v2"

var CmdHelp = &cli.Command{
	Name:      "help",
	Aliases:   []string{"h"},
	Usage:     "显示所有帮助信息或者显示单个指令的帮助信息(zenbox help setup)",
	ArgsUsage: "[command]",
	Action: func(c *cli.Context) error {
		args := c.Args()
		if args.Present() {
			return cli.ShowCommandHelp(c, args.First())
		}

		cli.ShowAppHelp(c)
		return nil
	},
}

func init() {
	cli.HelpFlag = &cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "显示帮助",
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "打印版本号",
	}

	cli.InitCompletionFlag = &cli.StringFlag{
		Name:  "init-completion",
		Usage: "生成命令行自动完成代码. 只支持 'bash' 和 'zsh'",
	}

	cli.AppHelpTemplate = `
{{.Name}} -- {{if .Usage}}{{.Usage}}{{end}}

用法:
   {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}} {{if .VisibleFlags}}[全局选项]{{end}}{{if .Commands}} 指令 [指令选项]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[参数...]{{end}}{{end}}{{if .Description}}

说明:
   {{.Description}}{{end}}{{if .VisibleCommands}}

指令列表:{{range .VisibleCategories}}{{if .Name}}
   {{.Name}}:{{end}}{{range .VisibleCommands}}
     {{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}{{end}}{{end}}{{if .VisibleFlags}}

全局选项:
   {{range $index, $option := .VisibleFlags}}{{if $index}}
   {{end}}{{$option}}{{end}}{{end}}{{if len .Authors}}

开发维护{{with $length := len .Authors}}{{if ne 1 $length}}{{end}}{{end}}:
   {{range $index, $author := .Authors}}{{if $index}}
   {{end}}{{$author}}{{end}}{{end}}

`
	cli.CommandHelpTemplate = `{{.HelpName}} - {{.Usage}}

用法:
   {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}}{{if .VisibleFlags}} [选项]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[参数...]{{end}}{{end}}{{if .Category}}

分类:
   {{.Category}}{{end}}{{if .Description}}

说明:
   {{.Description}}{{end}}{{if .VisibleFlags}}

选项:
   {{range .VisibleFlags}}{{.}}
   {{end}}{{end}}
`

	cli.SubcommandHelpTemplate = `NAME:
   {{.HelpName}} - {{if .Description}}{{.Description}}{{else}}{{.Usage}}{{end}}

USAGE:
   {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}} command{{if .VisibleFlags}} [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}

COMMANDS:{{range .VisibleCategories}}{{if .Name}}
   {{.Name}}:{{end}}{{range .VisibleCommands}}
     {{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}
{{end}}{{if .VisibleFlags}}
OPTIONS:
   {{range .VisibleFlags}}{{.}}
   {{end}}{{end}}
`
}
