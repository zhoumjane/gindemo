package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhoumjane/lfshook"
	"github.com/zhoumjane/logrus"
	"math"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetOutput(os.Stdout)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel: os.Stdout,
		logrus.FatalLevel: os.Stdout,
		logrus.DebugLevel: os.Stdout,
		logrus.WarnLevel: os.Stdout,
		logrus.ErrorLevel: os.Stdout,
		logrus.PanicLevel: os.Stdout,
	}
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		ForceColors:               false,
		DisableColors:             false,
		ForceQuote:                false,
		DisableQuote:              false,
		EnvironmentOverrideColors: false,
		DisableTimestamp:          false,
		FullTimestamp:             false,
		TimestampFormat:           "2006-01-02 15:04:05",
		DisableSorting:            false,
		SortingFunc:               nil,
		DisableLevelTruncation:    false,
		PadLevelText:              false,
		QuoteEmptyFields:          false,
		FieldMap:                  nil,
		CallerPrettyfier:          nil,
	})
	logger.AddHook(Hook)
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds()/1000000.0))))
		hostName, err := os.Hostname()
		if err != nil{
			hostName = "unknow"
		}
		statusCode := c.Writer.Status()
		clientIp := c.ClientIP()
		userAgent := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		method := c.Request.Method
		path := c.Request.RequestURI
		if dataSize < 0{
			dataSize = 0
		}
		entry := logger.WithFields(logrus.Fields{
			"hostName": hostName,
			"status": statusCode,
			"SpendTime": spendTime,
			"IP": clientIp,
			"Method": method,
			"Path": path,
			"DataSize": dataSize,
			"Agent": userAgent,
		})
		if len(c.Errors) > 0{
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500{
			entry.Error()
		}else if statusCode >= 400{
			entry.Warn()
		}else {
			entry.Info()
		}
	}
}