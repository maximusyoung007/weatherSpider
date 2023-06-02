package main

import (
	"github.com/jasonlvhit/gocron"
	"weatherSpider/business"
	"weatherSpider/conf"
	"weatherSpider/database"
	"weatherSpider/logu"
	"weatherSpider/structs"
)

func task() {
	log := logu.Logger
	log.Info("-----------------------------定时任务开始-----------------------------")
	gocron.Clear()
	var successList []structs.Area
	business.PreBusiness(&successList)
	database.InsertRow(successList)
	log.Info("-----------------------------定时任务结束-----------------------------")
}

func main() {
	log := logu.Logger
	log.Info("----------------------------程序启动--------------------------")
	conf.ConfigInit()
	logu.LogInit()
	log.Info("-----------------------------配置加载完成--------------------------------------")
	s := gocron.NewScheduler()
	s.Every(30).Minutes().Do(task)
	<-s.Start()
}
