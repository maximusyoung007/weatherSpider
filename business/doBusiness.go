package business

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"weatherSpider/client"
	"weatherSpider/convertData"
)

func DoBusiness() {
	mainUrl := "http://www.weather.com.cn/"
	cityJs := "https://j.i8tq.com/weather2020/search/city.js"
	citys := client.FetchString(cityJs)
	areaList := convertData.StructConvert(citys)
	doc, error := goquery.NewDocumentFromReader(client.Fetch(mainUrl))
	if error != nil {
		fmt.Print("获取goQuery内容失败")
	}
	//fmt.Println(doc.Find("p.category").Text())
	//fmt.Println(doc.Find("div.w_city").Find("dl").Find("dd").Find("a").Attr("href"))
	cityUrl, _ := doc.Find("div.w_city").Find("dl").Find("dd").Find("a[title~=上海]").Attr("href")
	fmt.Print(cityUrl)

	//101020100
	for i := 0; i < len(areaList); i++ {
		//strings.Replace()
	}
}
