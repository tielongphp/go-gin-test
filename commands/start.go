package commands

import (
	"fmt"
	"go-gin-test/context"
	"go-gin-test/server"
	"github.com/urfave/cli"
	"log"
	"github.com/joho/godotenv"
)

// Starts web server (user interface)
var StartCommand = cli.Command{
	Name:   "start",
	Usage:  "Starts web server",
	Flags:  startFlags,
	Action: startAction,
}

var startFlags = []cli.Flag{
	cli.IntFlag{
		Name:   "http-port, p",
		Usage:  "HTTP server port",
		Value:  8081,
		EnvVar: "GO_GIN_TEST_PORT",
	},
	cli.StringFlag{
		Name:   "http-host, i",
		Usage:  "HTTP server host",
		Value:  "",
		EnvVar: "GO_GIN_TEST_HOST",
	},
	cli.StringFlag{
		Name:   "http-mode, m",
		Usage:  "debug, release or test",
		Value:  "",
		EnvVar: "GO_GIN_TEST_MODE",
	},
}

func startAction(ctx *cli.Context) error {
	conf := context.NewConfig(ctx)

	//load ip file into memory
	/*fileName := conf.GetIpFile()
	myLogger := conf.GetLog()
	myLogger.Info("file  " + fileName)*/

	env := conf.GetEnv()

	if "" == env {
		env = "dev"
	}
	err:=godotenv.Load(".env." + env)
	if err != nil {
		log.Fatal("Error loading file .env."+ env)
	}

	if conf.HttpServerPort() < 1 {
		log.Fatal("Server port must be a positive integer")
	}

	fmt.Printf("Starting web server at %s:%d...\n", ctx.String("http-host"), ctx.Int("http-port"))

	server.Start(conf)

	fmt.Println("Done.")

	return nil
}
