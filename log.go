package util

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
		logDir = fmt.Sprintf("%s/logs", CWD())
	}
	log = logrus.New()
	//log.SetFormatter(&logrus.JSONFormatter{})
	name := fmt.Sprintf("%s/log", logDir)

	writer, _ := rotatelogs.New(
		name+"_%Y-%m-%d.log",
		rotatelogs.WithLinkName(name),             // 生成软链，指向最新日志文件
		rotatelogs.WithRotationTime(time.Hour*24), // 日志切割时间间隔
		rotatelogs.WithMaxAge(time.Hour*24*3),     //设置文件清理前的最长保存时间
	)
	log.SetOutput(writer)

	defer func() {
		_ = writer.Close()
	}()
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
