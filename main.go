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
	gocron.Clear()
	var successList []structs.Area
	business.PreBusiness(&successList)
	database.InsertRow(successList)
}

func main() {
	conf.ConfigInit()
	logu.LogInit()
	log := logu.Logger
	log.Info("-----------------------------定时任务开始-----------------------------")
	s := gocron.NewScheduler()
	s.Every(30).Minutes().Do(task)
	<-s.Start()
	log.Info("-----------------------------定时任务结束-----------------------------")
}
