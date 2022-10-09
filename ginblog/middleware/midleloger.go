package middleware

import (
	"fmt"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
	retalog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"path"
	"time"
)

// 日志记录到文件
func LoggerToFile() gin.HandlerFunc {

	logFilePath := utils.LogFile
	logFileName := utils.LogName

	fmt.Println("path = ", path.Join(logFilePath, logFileName))

	//日志文件
	fileName := "./log/logger.txt" //path.Join(logFilePath, logFileName)

	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	//实例化
	logger := logrus.New()

	//设置输出
	logger.Out = src

	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	//设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{})

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}

func Logger() gin.HandlerFunc {
	logFilePath := utils.LogFile
	logFileName := utils.LogName

	fmt.Println("path = ", path.Join(logFilePath, logFileName))

	//日志文件
	fileName := "log/log" //path.Join(logFilePath, logFileName)
	linkFile := "link.log"

	//写入文件
	//src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)  //os.ModeAppend

	src, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755) //os.ModeAppend
	if err != nil {
		fmt.Println("err", err)
	}

	log := logrus.New()
	//设置输出
	log.Out = src

	log.SetLevel(logrus.DebugLevel)
	//日志分割

	writer, _ := retalog.New(
		fileName+"_%Y%m%d.log",
		//生成软链 指向最新日志文件
		retalog.WithLinkName(linkFile),
		//最大保存时间
		retalog.WithMaxAge(7*24*time.Hour),
		//日志切割时间
		retalog.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  writer,
		logrus.FatalLevel: writer,
		logrus.DebugLevel: writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.PanicLevel: writer,
		logrus.TraceLevel: writer,
	}

	//加入HooK
	hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05", //时间一定要是这个，否则无法Hook
	})

	//加入hook
	log.AddHook(hook)

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds()/1000000.0))))
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown hostname"
		}

		statusCode := c.Writer.Status()
		clientIp := c.ClientIP()
		userAgent := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}

		method := c.Request.Method
		path := c.Request.RequestURI

		entry := log.WithFields(logrus.Fields{

			"hostName":  hostName,
			"status":    statusCode,
			"spendTime": spendTime,
			"ip":        clientIp,
			"userAgent": userAgent,
			"method":    method,
			"path":      path,
		})

		if len(c.Errors) > 0 { //gin 框架 内部报错  记录gin 内部错误
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}

		//HTTP response status HTTP 记录
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}

	}
}

// 日志记录到 MongoDB
func LoggerToMongo() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// 日志记录到 ES
func LoggerToES() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// 日志记录到 MQ
func LoggerToMQ() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
