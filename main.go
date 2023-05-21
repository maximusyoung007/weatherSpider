package main

import (
	"fmt"
	"weatherSpider/business"
	"weatherSpider/structs"
)

func main() {
	fmt.Println("hello world")
	//client.Connect("http://www.weather.com.cn/")
	var successList []structs.Area
	println(business.PreBusiness(&successList))
	for _, v := range successList {
		fmt.Println(v.NameCN)
	}
}
