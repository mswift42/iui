package main

import (
	"github.com/mswift42/iui/cl"
	"os"
)

func main() {
	app := cl.InitCli()
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}

}
