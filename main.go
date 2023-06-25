package main

import (
	"github.com/jasonlvhit/gocron"
	"github.com/sirupsen/logrus"
	"weatherSpider/business"
	"weatherSpider/conf"
	"weatherSpider/database"
	"weatherSpider/logu"
	"weatherSpider/mail"
	"weatherSpider/structs"
)

var log = &logu.Logger
var co = &conf.Conf

func task() {
	(*log).Info("-----------------------------定时任务开始-----------------------------")
	//gocron.Clear()
	var successList []structs.Area
	business.PreBusiness(&successList)
	database.InsertRow(successList)
	(*log).Info("-----------------------------定时任务结束-----------------------------")
}

func main() {
	conf.ConfigInit()
	logu.LogInit()
	(*log).WithFields(logrus.Fields{"日志文件名": co.Log.FileName, "日志文件路径": co.Log.FilePath}).
		Info("-----------------------------配置加载完成--------------------------------------")

	s := gocron.NewScheduler()
	s.Every(60).Minutes().Do(task)
	s.Every(1).Day().At("23:45").Do(mail.SendEmail)
	<-s.Start()

}
