package business

import (
	"github.com/PuerkitoBio/goquery"
	"weatherSpider/convertData"
)

func doBusiness() {
	cityUrl := "https://j.i8tq.com/weather2020/search/city.js"
	areaList := convertData.StructConvert(cityUrl)
	doc, error := goquery.
}
