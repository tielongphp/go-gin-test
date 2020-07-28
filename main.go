package main

import (
	"os"

	"github.com/urfave/cli"

	"go-gin-test/commands"
)

var version = "development"

func main() {
	app := cli.NewApp()
	app.Name = "go-gin-test"
	app.Usage = ""
	app.Version = version
	app.Copyright = "(c) 2020 go-gin-test"
	app.EnableBashCompletion = true
	app.Flags = commands.GlobalFlags

	app.Commands = []cli.Command{
		commands.ConfigCommand,
		commands.StartCommand,
	}

	app.Run(os.Args)
}
