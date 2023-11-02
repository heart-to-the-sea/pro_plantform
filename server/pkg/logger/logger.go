package logger

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.SetReportCaller(false)
}

func Write(msg string, filename string) {
	setOutPutFile(logrus.InfoLevel, filename)
	logrus.Info(msg)
}

func Info(fields logrus.Fields, args ...interface{}) {
	setOutPutFile((logrus.FatalLevel), "info")
	logrus.WithFields(fields).Info(args)
}

func Recover(c *gin.Context) {
	defer func() {
		filePath := "./runtime/log"
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			err = os.MkdirAll(filePath, 0777)
			if err != nil {
				panic(fmt.Errorf("create log dir '%s' error: %s", filePath, err))
			}
		}

		timeStr := time.Now().Format("2006-01-02")
		fileName := path.Join(filePath, "error"+timeStr+".log")

		f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		timeFileStr := time.Now().Format("2006-01-02 15:04:05")
		f.WriteString("panic error time:" + timeFileStr + "\n")
		f.WriteString(fmt.Sprintf("%v", err) + "\n")
		f.WriteString("stacktrace from panic:" + string(debug.Stack()) + "\n")
		f.Close()
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  fmt.Sprintf("%v", err),
		})
		c.Abort()
	}()
	c.Next()
}

/** 输出到指定文件目录下 */
func setOutPutFile(level logrus.Level, logName string) {
	filePath := "./runtime/log"
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		err = os.MkdirAll(filePath, 0777)
		if err != nil {
			panic(fmt.Errorf("create log dir '%s' error: %s", filePath, err))
		}
	}

	timeStr := time.Now().Format("2006-01-02")
	fileName := path.Join(filePath, logName+timeStr+".log")

	var err error
	os.Stderr, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("open log file err", err)
	}
	logrus.SetOutput(os.Stderr)
	logrus.SetLevel(level)

}
func LoggerToFile() gin.LoggerConfig {
	filePath := "./runtime/log"
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		err = os.MkdirAll(filePath, 0777)
		if err != nil {
			panic(fmt.Errorf("create log dir '%s' error: %s", filePath, err))
		}
	}

	timeStr := time.Now().Format("2006-01-02")
	fileName := path.Join(filePath, "gin"+timeStr+".log")

	os.Stderr, _ = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	var conf = gin.LoggerConfig{
		Formatter: func(params gin.LogFormatterParams) string {
			return fmt.Sprintf("%s - %s \"%s %s %s %d %s \"%s \" %s \"\n",
				params.TimeStamp.Format("2006-01-02 15:04:05"),
				params.ClientIP,
				params.Method,
				params.Path,
				params.Request.Proto,
				params.StatusCode,
				params.Latency,
				params.Request.UserAgent(),
				params.ErrorMessage,
			)
		},
		Output: io.MultiWriter(os.Stdout, os.Stderr),
	}
	return conf
}
