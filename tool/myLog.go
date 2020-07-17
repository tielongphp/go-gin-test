package tool

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type MyLog struct {
	Logger *logrus.Logger
}

func NewMyLog() *MyLog {
	log := &MyLog{}
	log.setMyLogger()
	return log
}

type ErrorApiRequest struct {
	ServiceStart     time.Time              `json:"serviceStart"`
	Elapsed          float64                `json:"elapsed"`
	ApiRequestUri    string                 `json:"apiRequestUri"`
	Arguments        map[string]interface{} `json:"arguments"`
	Status           int                    `json:"status"`
	ApiRequestMethod string                 `json:"apiRequestMethod"`
	ServiceEnd       time.Time              `json:"serviceEnd"`
	Result           string                 `json:"result"`
}

func (myLog *MyLog) Info(desc string, c *gin.Context) {
	//endTime := time.Now().Format("2006-01-02 15:04:05")
	hostName, _ := os.Hostname()
	request := c.Request
	traceId := request.Header.Get("X-Ca-Traceid")
	uuidV := uuid.NewV4()
	if traceId == "" {
		request.Header.Set("X-Ca-Traceid", uuidV.String())
		traceId = uuidV.String()
	}

	descMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(desc), &descMap)
	apiRequestForApiCall := make(map[string]interface{})
	apiRequestForApiCall["desc"] = "###" + desc + "###"
	end := time.Now()
	start, _ := c.Get("go-gin-test-start")

	if err == nil {
		//namelookupTime, ok1 := descMap["namelookupTime"]
		elapsed, ok2 := descMap["elapsed"]
		if ok2 {
			//apiRequestForApiCall["namelookupTime"] = namelookupTime
			apiRequestForApiCall["elapsed"] = elapsed
			myLog.Logger.WithFields(logrus.Fields{
				"clientIp":      c.ClientIP(),
				"hostAddress":   GetHostAddress(),
				"hostName":      hostName,
				"logLevel":      "info",
				"logType":       "run",
				"requestMethod": request.Method,
				"requestRoute":  request.URL.Path,
				"requestUri":    request.RequestURI,
				"serviceEnd":    end,
				"serviceStart":  start,
				"traceId":       traceId,
				"apiRequest":    apiRequestForApiCall,
				"context":       request.RequestURI,
			}).Info("info log")
		}
	} else {
		myLog.Logger.WithFields(logrus.Fields{
			"clientIp":      c.ClientIP(),
			"hostAddress":   GetHostAddress(),
			"hostName":      hostName,
			"logLevel":      "info",
			"logType":       "run",
			"requestMethod": request.Method,
			"requestRoute":  request.URL.Path,
			"requestUri":    request.RequestURI,
			"serviceEnd":    end,
			"serviceStart":  start,
			"traceId":       traceId,
			"apiRequest":    apiRequestForApiCall,
			"context":       request.RequestURI,
		}).Info("info log")
	}
}

func (myLog *MyLog) Error(desc string, c *gin.Context) {
	hostName, _ := os.Hostname()
	request := c.Request
	traceId := request.Header.Get("X-Ca-Traceid")
	uuidV := uuid.NewV4()
	if traceId == "" {
		request.Header.Set("X-Ca-Traceid", uuidV.String())
		traceId = uuidV.String()
	}
	end := time.Now()
	start, _ := c.Get("go-gin-test-start")
	apiRequestForApiCall := make(map[string]interface{})
	apiRequestForApiCall["desc"] = "###" + desc + "###"
	myLog.Logger.WithFields(logrus.Fields{
		"clientIp":      c.ClientIP(),
		"hostAddress":   GetHostAddress(),
		"hostName":      hostName,
		"logLevel":      "emerg",
		"logType":       "error",
		"requestMethod": request.Method,
		"requestRoute":  request.URL.Path,
		"requestUri":    request.RequestURI,
		"serviceEnd":    end,
		"serviceStart":  start,
		"traceId":       traceId,
		"apiRequest":    apiRequestForApiCall,
		"context":       request.RequestURI,
	}).Info("error log")

}

func (myLog *MyLog) setMyLogger() {
	var log = logrus.New()
	log.Formatter = new(logrus.JSONFormatter)

	log.Level = logrus.TraceLevel
	log.Out = os.Stdout

	defer func() {
		err := recover()
		if err != nil {
			entry := err.(*logrus.Entry)
			log.WithFields(logrus.Fields{
				"omg":         true,
				"err_animal":  entry.Data["animal"],
				"err_size":    entry.Data["size"],
				"err_level":   entry.Level,
				"err_message": entry.Message,
				"number":      100,
			}).Error("The ice breaks!") // or use Fatal() to force the process to exit with a nonzero code
		}
	}()

	myLog.Logger = log
}
