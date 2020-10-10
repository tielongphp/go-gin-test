package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"go-gin-test/context"
	"go-gin-test/tool"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"
)

func Logger(conf *context.Config) gin.HandlerFunc {
	logClient := conf.GetMyLogger().Logger
	//禁止logrus的输出
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("Open Src File err", err)
	}
	logClient.SetOutput(src)
	logClient.SetLevel(log.DebugLevel)
	apiLogPath := conf.GetLogDir()
	baseLogPath := path.Join(os.Getenv("APP_PATH"), apiLogPath, "log")
	writer, err := rotatelogs.New(
		baseLogPath+".%Y-%m-%d.log",
		rotatelogs.WithLinkName(baseLogPath),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	if err != nil {
		log.Fatal("Init log failed, err:", err)
	}

	logClient.SetOutput(writer)
	logClient.SetLevel(log.InfoLevel)

	//logClient := conf.GetMyLogger().Logger
	//
	////禁止logrus的输出
	//src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	//if err != nil {
	//	fmt.Println("err", err)
	//}
	//logClient.Out = src
	//logClient.SetLevel(logrus.DebugLevel)
	//
	//apiLogPath := conf.GetLogDir()
	//fmt.Println("apiLogPath" + "  " + apiLogPath)
	//logWriter, err := rotatelogs.New(
	//	apiLogPath+".%Y-%m-%d.log",
	//	rotatelogs.WithLinkName(apiLogPath),       // 生成软链，指向最新日志文件
	//	rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
	//	rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	//)
	//writeMap := lfshook.WriterMap{
	//	logrus.InfoLevel:  logWriter,
	//	logrus.FatalLevel: logWriter,
	//	logrus.ErrorLevel: logWriter,
	//	logrus.PanicLevel: logWriter,
	//	logrus.DebugLevel: logWriter,
	//
	//}
	//lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{})
	//logClient.AddHook(lfHook)

	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()
		c.Set("go-gin-test-start", start)
		// 处理请求
		c.Next()
		// 结束时间
		end := time.Now()
		//执行时间
		latency := end.Sub(start)

		path := c.Request.URL.Path

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()

		hostName, _ := os.Hostname()
		traceId := c.Request.Header.Get("X-Ca-Traceid")
		uuidV := uuid.NewV4()
		if traceId == "" {
			c.Request.Header.Set("X-Ca-Traceid", uuidV.String())
			traceId = uuidV.String()
		}

		elapsed := tool.FormatElapse(latency.Nanoseconds())

		arguments := ""
		if strings.ToUpper(method) == "POST" {
			argument, _ := ioutil.ReadAll(c.Request.Body)
			arguments = string(argument)
		}

		logClient.WithFields(log.Fields{
			"arguments":     arguments,
			"clientIp":      clientIP,
			"elapsed":       elapsed,
			"hostAddress":   tool.GetHostAddress(),
			"hostName":      hostName,
			"logLevel":      "access",
			"logType":       "access",
			"referer":       c.Request.Referer(),
			"requestHost":   c.Request.Host,
			"requestMethod": method,
			"requestRoute":  path,
			"requestUri":    c.Request.RequestURI,
			"result":        "",
			"serviceEnd":    end,
			"serviceStart":  start,
			"status":        statusCode,
			"traceId":       traceId,
			"context":       c.Request.RequestURI,
		}).Info("access log")
	}
}
