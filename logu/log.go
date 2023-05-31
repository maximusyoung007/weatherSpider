package logu

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
	"weatherSpider/conf"
)

var Logger *logrus.Logger

func LogInit() {
	Logger = logrus.New()
	//以json格式输出
	Logger.Formatter = &logrus.JSONFormatter{}
	Logger.SetLevel(logrus.DebugLevel)
	//显示代码在什么位置打印
	Logger.SetReportCaller(true)
	path := conf.Conf.Log.FilePath + conf.Conf.Log.FileName
	writer, err := rotatelogs.New(
		path+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(time.Duration(72)*time.Hour),
		rotatelogs.WithRotationSize(1024),
	)
	//同时写到终端和文件
	fileAndStdoutWriter := io.MultiWriter(os.Stdout, writer)
	if err == nil {
		Logger.SetOutput(fileAndStdoutWriter)
	} else {
		Logger.Info("打开日志文件失败，默认输出到stdErr")
	}
}
