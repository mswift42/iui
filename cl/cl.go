package cl

import (
	"github.com/urfave/cli"
)

func InitCli() *cli.App {
	app := cli.NewApp()
	app.Setup()
	app.Name = "iui"
	app.Author = "Martin Haesler"
	app.Version = "0.0.1"
	app.Copyright = "2019 (c) Martin Haesler"
	app.Usage = "generate custom ui plugins for IntelliJ editors > 191"
	cli.AppHelpTemplate = `NAME:
   {{.Name}} - {{.Usage}}
	
USAGE:
   {{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}{{if .Commands}} command 
	[command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments. ..]{{end}}
AUTHOR: {{.Author}}
   {{if .Commands}}
COMMANDS:
{{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
GLOBAL OPTIONS:
   {{range .VisibleFlags}}{{.}}
   {{end}}{{end}}
COPYRIGHT:
   {{.Copyright}}
VERSION:
   {{.Version}}
`
	return app
}
