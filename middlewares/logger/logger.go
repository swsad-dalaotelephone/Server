package logger

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {
	logClient := log.New()
	//禁止logrus的输出，只通过hook输出
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	logClient.SetOutput(src)
	// set writer to cut log file
	GOPATH := os.Getenv("GOPATH")
	logPath := GOPATH + "/src/github.com/swsad-dalaotelephone/Server/storage/logs"
	fileName := path.Join(logPath, "api.log")
	logWriter, err := rotatelogs.New(
		fileName+".%Y-%m-%d-%H-%M.log",
		rotatelogs.WithLinkName(fileName),         // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	// set hook
	writeMap := lfshook.WriterMap{
		log.DebugLevel: logWriter, // 为不同级别设置不同的输出目的
		log.InfoLevel:  logWriter,
		log.WarnLevel:  logWriter,
		log.ErrorLevel: logWriter,
		log.FatalLevel: logWriter,
		log.PanicLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &log.JSONFormatter{})
	logClient.AddHook(lfHook)

	return func(c *gin.Context) {
		// latency
		start := time.Now()
		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		// msg
		msg := "Request"
		if len(c.Errors) > 0 {
			msg = c.Errors.String()
		}
		//path
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}

		contextLogger := logClient.WithFields(log.Fields{
			"status":  c.Writer.Status(),
			"method":  c.Request.Method,
			"path":    path,
			"ip":      c.ClientIP(),
			"latency": latency,
		})
		switch {
		case c.Writer.Status() >= http.StatusBadRequest && c.Writer.Status() < http.StatusInternalServerError:
			contextLogger.Warn(msg)
			break
		case c.Writer.Status() >= http.StatusInternalServerError:
			contextLogger.Error(msg)
			break
		default:
			contextLogger.Info(msg)
			break
		}
	}
}
