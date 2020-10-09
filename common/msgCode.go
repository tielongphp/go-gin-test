package common

import (
	"github.com/gin-gonic/gin"

	"go-gin-test/response"
)

const TheOther = "THE_OTHER"
const ParamError = "PARAM_ERROR"
const SystemError = "SYSTEM_ERROR"
const TOKEN_EXPIRE = "TOKEN_EXPIRE"
const DATA_NOT_FIND = "DATA_NOT_FIND"

type MsgCode struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func getAllAllCodeMsg() map[string]MsgCode {
	codeMsgMap := make(map[string]MsgCode)
	codeMsgMap[ParamError] = MsgCode{Code: 40000, Msg: "参数错误/参数为空"}
	codeMsgMap[TheOther] = MsgCode{Code: 99999, Msg: "其他错误"}
	codeMsgMap[SystemError] = MsgCode{Code: 99998, Msg: "服务系统异常"}
	codeMsgMap[TOKEN_EXPIRE] = MsgCode{Code: 50001, Msg: "token失效或者token不存在"}
	codeMsgMap[DATA_NOT_FIND] = MsgCode{Code: 40002, Msg: "数据不存在"}

	return codeMsgMap
}

func GetCodeMsg(msg string, c *gin.Context) {
	codeMsgMap := getAllAllCodeMsg()
	if data, ok := codeMsgMap[msg]; ok {
		response.ResultNotWithData(data.Code, data.Msg, c)
	} else {
		response.ResultNotWithData(codeMsgMap[SystemError].Code, codeMsgMap[SystemError].Msg, c)
	}
}

func returnCodeMsg(msg string, c *gin.Context) response.NotWithData {
	codeMsgMap := getAllAllCodeMsg()
	if data, ok := codeMsgMap[msg]; ok {
		return response.NotWithData{
			Code: data.Code,
			Msg:  data.Msg,
		}
	} else {
		return response.NotWithData{
			Code: codeMsgMap[SystemError].Code,
			Msg:  codeMsgMap[SystemError].Msg,
		}
	}
}
