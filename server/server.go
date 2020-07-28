package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-test/context"
	"go-gin-test/models"
)

// Start the REST API server using the configuration provided
func Start(conf *context.Config) {
	if conf.HttpServerMode() != "" {
		gin.SetMode(conf.HttpServerMode())
	} else if conf.Debug() == false {
		gin.SetMode(gin.ReleaseMode)
	}

	/*logFile := conf.LogFilePath()
	gin.DefaultWriter = io.MultiWriter(logFile)*/
	app := gin.Default()

	// route
	registerRoutes(app, conf)
	// DB
	models.Init(conf)

	app.Run(fmt.Sprintf("%s:%d", conf.HttpServerHost(), conf.HttpServerPort()))
}
