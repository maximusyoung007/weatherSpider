package business

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"weatherSpider/chromeDp"
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
	oldAreaId := ""
	for i := 0; i < len(areaList); i++ {
		//strings.Replace()
		area := areaList[i]
		if i == 0 {
			cityUrl = strings.Replace(cityUrl, "101020100", area.AreaId, 1)
		} else {
			cityUrl = strings.Replace(cityUrl, oldAreaId, area.AreaId, 1)
		}
		oldAreaId = area.AreaId
		//docArea, errArea := goquery.NewDocumentFromReader(client.Fetch(mainUrl))
		//if errArea != nil {
		//	fmt.Print("获取地区天气内容失败")
		//}
		if i == 0 {
			sa, _ := chromeDp.GetHttpHtmlContent(cityUrl, "#today > div.t > div > div.zs.pol > span > a", "document.querySelector(\"body\")")
			fmt.Println(sa)
		}
	}
}
