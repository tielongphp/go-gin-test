package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int         `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type NotWithData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

const (
	ERROR   = 99999
	SUCCESS = 100000
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}
func ResultNotWithData(code int, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, NotWithData{
		code,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "成功", c)
}

func OkWithMsg(msg string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, msg, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "成功", c)
}

func OkDetailed(data interface{}, msg string, c *gin.Context) {
	Result(SUCCESS, data, msg, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "失败", c)
}

func FailWithMsg(msg string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, msg, c)
}

func FailWithDetailed(code int, data interface{}, msg string, c *gin.Context) {
	Result(code, data, msg, c)
}
