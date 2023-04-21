package main

import (
	"fmt"
	"weatherSpider/client"
	"weatherSpider/convertData"
)

func main() {
	fmt.Println("hello world")
	//client.Connect("http://www.weather.com.cn/")
	s := client.Connect("https://j.i8tq.com/weather2020/search/city.js")
	convertData.StructConvert(s)
}
