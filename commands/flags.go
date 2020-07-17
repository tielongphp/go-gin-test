package commands

import (
	"github.com/urfave/cli"
	"path"
	"os"
)

// Global CLI flags
var GlobalFlags = []cli.Flag{
	cli.StringFlag{
		Name:   "env",
		Usage:  "use -env=dev/stage/prod",
		EnvVar: "GO_GIN_TEST_ENV",
	},
	cli.BoolFlag{
		Name:   "debug",
		Usage:  "run in debug mode",
		EnvVar: "GO_GIN_TEST_DEBUG",
	},
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
	cli.StringFlag{
		Name:   "log-dir, ld",
		Usage:  "log dir",
		Value:  path.Join(os.Getenv("APP_PATH"),"logs/"),
		EnvVar: "GO_GIN_TEST_LOG_DIR",
	},
}
