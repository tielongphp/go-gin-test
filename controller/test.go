package controller

import (
	"github.com/gin-gonic/gin"
	"go-gin-test/common"
	"go-gin-test/context"
)

func Test(router *gin.RouterGroup, conf *context.Config) {
	router.GET("/a", func(c *gin.Context) {

		// todo
		common.FormatResponseWithoutData(c, 10000, "成功")
	})
}
