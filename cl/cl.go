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
	return app
}
