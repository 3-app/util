package log

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var log *logrus.Logger

func Init() {
	logDir := os.Getenv("LOG_DIR")
	if logDir == "" {
		logDir = "/data"
	}
	//测试记录
	log = logrus.New()
	//log.Hooks.Add(NewContextHook())  //增加行号hook
	//log.SetFormatter(&logrus.JSONFormatter{})
	name := fmt.Sprintf("%s/log", logDir)

	writer, _ := rotatelogs.New(
		name+"_%Y-%m-%d-%H.log",
		rotatelogs.WithLinkName(name),          // 生成软链，指向最新日志文件
		rotatelogs.WithRotationTime(time.Hour), // 日志切割时间间隔
		rotatelogs.WithMaxAge(time.Hour),       //设置文件清理前的最长保存时间
	)
	log.SetOutput(writer)

	defer func() {
		_ = writer.Close()
	}()
}

func GetLogger() *logrus.Logger {
	if log == nil {
		Init()
	}
	return log
}

func Err(err error, msg string) {
	fmt.Println(fmt.Sprintf("%s:%s", msg, err.Error()))
	if err != nil {
		log.Fatalf("%s: %s", msg, err.Error())
	}
}

func Info(msg string) {
	log.Infof(msg)
}

func Warning(msg string) {
	fmt.Println(msg)
	log.Warningf(msg)
}
