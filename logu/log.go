package logu

import (
	"github.com/sirupsen/logrus"
	"os"
)

var logger *logrus.Logger

func logInit() {
	logger = logrus.New()
	//以json格式输出
	logger.Formatter = &logrus.JSONFormatter{}
	logger.SetLevel(logrus.DebugLevel)
	//显示代码在什么位置打印
	logger.SetReportCaller(true)
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.SetOutput(file)
	} else {
		logger.Info("打开日志文件失败，默认输出到stdErr")
	}
}
