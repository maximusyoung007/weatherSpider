package main

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"weatherSpider/business"
	"weatherSpider/database"
	"weatherSpider/structs"
)

func task() {
	fmt.Println("task begin")
	gocron.Clear()
	var successList []structs.Area
	business.PreBusiness(&successList)
	database.InsertRow(successList)
}

func main() {
	s := gocron.NewScheduler()
	s.Every(30).Minutes().Do(task)
	<-s.Start()
	//gocron.Every(1).Hours().Do(task)
}
