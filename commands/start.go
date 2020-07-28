package commands

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"

	"go-gin-test/context"
	"go-gin-test/server"
)

// Starts web server (user interface)
var StartCommand = cli.Command{
	Name:   "start",
	Usage:  "Starts web server",
	Flags:  GlobalFlags,
	Action: startAction,
}

//var startFlags = GlobalFlags
//cli.StringFlag{
//	Name:   "env, e",
//	Usage:  "env",
//	Value:  "dev",
//	EnvVar: "GO_GIN_TEST_HOST",
//},
//cli.IntFlag{
//	Name:   "http-port, p",
//	Usage:  "HTTP server port",
//	Value:  8081,
//	EnvVar: "GO_GIN_TEST_PORT",
//},
//cli.StringFlag{
//	Name:   "http-host, i",
//	Usage:  "HTTP server host",
//	Value:  "",
//	EnvVar: "GO_GIN_TEST_HOST",
//},
//cli.StringFlag{
//	Name:   "http-mode, m",
//	Usage:  "debug, release or test",
//	Value:  "",
//	EnvVar: "GO_GIN_TEST_MODE",
//},
//}

//var Cfg *ini.File

func startAction(ctx *cli.Context) error {
	conf := context.NewConfig(ctx)

	//load ip file into memory
	/*fileName := conf.GetIpFile()
	myLogger := conf.GetLog()
	myLogger.Info("file  " + fileName)*/

	os.Setenv("APP_PATH", conf.GetCfg().Section("").Key("APP_PATH").String())
	os.Setenv("ENV", conf.GetEnv())
	if conf.HttpServerPort() < 1 {
		log.Fatal("Server port must be a positive integer")
	}

	fmt.Printf("Starting web server at %s:%d...\n", ctx.String("http-host"), ctx.Int("http-port"))

	server.Start(conf)

	fmt.Println("Done.")

	return nil
}
