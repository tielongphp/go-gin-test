package common

import (
	"github.com/gin-gonic/gin"
)

type result struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

type multiResult struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	DataList    interface{} `json:"dataList"`
}

type resultWithoutDate struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

func FormatResponse(c *gin.Context, code int, msg string, data interface{}) {
	if msg == "" {
		msg = "成功"
	}
	c.JSON(200, &result{
		code,
		msg,
		data,
	})
}

func MultiFormatResponse(c *gin.Context, code int, msg string, data interface{}) {
	if msg == "" {
		msg = "成功"
	}
	c.JSON(200, &multiResult{
		code,
		msg,
		data,
	})
}

func FormatResponseWithoutData(c *gin.Context, code int, msg string) {
	c.JSON(200, &resultWithoutDate{
		code,
		msg,
	})
	return
}
