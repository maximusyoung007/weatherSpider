package main

import (
	"github.com/jasonlvhit/gocron"
	"github.com/sirupsen/logrus"
	"weatherSpider/conf"
	"weatherSpider/logu"
	"weatherSpider/mail"
)

var log = &logu.Logger
var co = &conf.Conf

func task() {
	(*log).Info("-----------------------------定时任务开始-----------------------------")
	//gocron.Clear()
	//var successList []structs.Area
	//business.PreBusiness(&successList)
	//database.InsertRow(successList)
	(*log).Info("-----------------------------定时任务结束-----------------------------")
}

func main() {
	conf.ConfigInit()
	logu.LogInit()
	(*log).WithFields(logrus.Fields{"日志文件名": co.Log.FileName, "日志文件路径": co.Log.FilePath}).
		Info("-----------------------------配置加载完成--------------------------------------")
	s := gocron.NewScheduler()
	//mail.SendEmail()
	s.Every(5).Minutes().Do(mail.SendEmail)
	s.Every(4).Minutes().Do(task)
	<-s.Start()
}
