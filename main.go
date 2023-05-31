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
	//log = logu.Logger
	//logu.Logger.Info("-----------------------------初始化完成！-----------------------------")
	//log.Info("-----------------------------定时任务开始-----------------------------")
	s := gocron.NewScheduler()
	s.Every(30).Seconds().Do(task)
	<-s.Start()
	//log.Info("-----------------------------定时任务结束-----------------------------")
	//gocron.Every(1).Hours().Do(task)
}
