package cl

import (
	"github.com/mswift42/iui/ui"
	"io/ioutil"
	"os"

	"github.com/urfave/cli"
)

func loadFile(fp string) ([]byte, error) {
	file, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(file)
}

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
	app.Commands = []cli.Command{
		{
			Name:      "generate",
			Aliases:   []string{"g", "gen", "generate"},
			Usage:     "Generate ui theme from supplied editor theme or json color map and a ui theme template.",
			HelpName:  "generate",
			ArgsUsage: "[theme/color map] [ui theme template]",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "json, j",
					Usage: "Use if you are using the json template to generate a ui plugin.",
				},
			},
			Action: func(c *cli.Context) error {
				if len(c.Args()) < 2 {
					return cli.NewExitError("Generate Format is <theme scheme> <ui template", 1)
				}
				scheme := c.Args().Get(0)
				templpath := c.Args().Get(1)
				if c.Bool("json") {
					return ui.GenerateThemeFromJson(scheme, templpath)
				} else {
					return ui.GenerateTheme(scheme, templpath)
				}
				return nil
			},
		},
	}
	return app
}
