package server

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"go-gin-test/context"
	"go-gin-test/global"
	"go-gin-test/initialize"
)

// Start the REST API server using the configuration provided
func Start(conf *context.Config) {
	global.CTX_CONFIG = conf
	if conf.HttpServerMode() != "" {
		gin.SetMode(conf.HttpServerMode())
	} else if conf.Debug() == false {
		gin.SetMode(gin.ReleaseMode)
	}

	/*logFile := conf.LogFilePath()
	gin.DefaultWriter = io.MultiWriter(logFile)*/
	app := gin.Default()
	//初始化 utils
	initialize.UtilsInit()
	//初始化 DB
	initialize.MysqlInit()

	// 初始化redis
	initialize.RedisInit()
	// 初始化总路由
	initialize.RoutersInit(app, conf)

	//defer model.CloseDB()
	app.Run(fmt.Sprintf("%s:%d", conf.HttpServerHost(), conf.HttpServerPort()))
}
