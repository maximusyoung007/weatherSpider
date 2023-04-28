package main

import (
	"fmt"
	"weatherSpider/business"
)

func main() {
	fmt.Println("hello world")
	//client.Connect("http://www.weather.com.cn/")
	business.DoBusiness()
}
