package cl

import (
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
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
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "scheme, s",
			Value: "xml",
			Usage: "theme file to parse",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:      "generate",
			Aliases:   []string{"g", "gen", "generate"},
			Usage:     "Generate ui theme from supplied editor theme or json color map.",
			HelpName:  "generate",
			ArgsUsage: "[theme/color map]",
			Action: func(c *cli.Context) error {
				fmt.Println(c.Args().Get(0))
				return nil
			},
		},
	}
	return app
}
