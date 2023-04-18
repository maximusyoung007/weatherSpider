package main

import (
	"fmt"
	"weatherSpider/client"
)

func main() {
	fmt.Println("hello world")
	client.Connect("https://news.baidu.com/")
}
