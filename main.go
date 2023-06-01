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
	conf.ConfigInit()
	logu.LogInit()
	s := gocron.NewScheduler()
	s.Every(30).Seconds().Do(task)
	<-s.Start()
}
