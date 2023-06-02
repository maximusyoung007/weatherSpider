package business

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"weatherSpider/chromeDp"
	"weatherSpider/client"
	"weatherSpider/convertData"
	"weatherSpider/logu"
	"weatherSpider/structs"
)

/*
*
两个areaList,一个是处理完成的，一个是异常的处理
*/
var log = &logu.Logger

func DoBusiness(areaSuccess *[]structs.Area, areaList []structs.Area) string {
	mainUrl := "http://www.weather.com.cn/"
	processingList := make([]structs.Area, 0)
	for i := 0; i < len(areaList); i++ {
		if areaList[i].AreaId == "101010100" || areaList[i].AreaId == "101020100" || areaList[i].AreaId == "101270101" ||
			areaList[i].AreaId == "101210101" || areaList[i].AreaId == "101190101" || areaList[i].AreaId == "101030100" ||
			areaList[i].AreaId == "101280601" || areaList[i].AreaId == "101040100" || areaList[i].AreaId == "101110101" ||
			areaList[i].AreaId == "101280101" || areaList[i].AreaId == "101120201" || areaList[i].AreaId == "101200101" ||
			areaList[i].AreaId == "101190401" {
			processingList = append(processingList, areaList[i])
		}
	}
	areaListErr := make([]structs.Area, 0)
	doc, error := goquery.NewDocumentFromReader(client.Fetch(mainUrl))
	if error != nil {
		(*log).WithFields(logrus.Fields{"error": error}).Info("获取goQuery内容失败")
	}
	cityUrl, _ := doc.Find("div.w_city").Find("dl").Find("dd").Find("a[title~=上海]").Attr("href")

	oldAreaId := ""
	for i := 0; i < len(processingList); i++ {
		area := processingList[i]
		if i == 0 {
			cityUrl = strings.Replace(cityUrl, "101020100", area.AreaId, 1)
		} else {
			cityUrl = strings.Replace(cityUrl, oldAreaId, area.AreaId, 1)
		}
		oldAreaId = area.AreaId
		//每隔一秒查询一个
		//time.Sleep(1e9)
		htmlContent, chromeErr := chromeDp.GetHttpHtmlContent(cityUrl, "#today > div.t > div > div.zs.pol > span > a", "document.querySelector(\"body\")")
		if chromeErr != nil {
			areaListErr = append(areaListErr, area)
			continue
		}
		dom, _ := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
		ac, _ := dom.Find("div.zs.pol").Find("span").Find("a").Html()
		for i := 0; i < len(ac); i++ {
			if ac[i] >= '0' && ac[i] <= '9' {
				continue
			} else {
				ac = ac[:i]
			}
		}
		acn, _ := strconv.Atoi(ac)
		area.AirCondition = acn
		*areaSuccess = append(*areaSuccess, area)
	}
	if len(areaListErr) > 0 {
		return DoBusiness(areaSuccess, areaListErr)
	} else {
		return "success"
	}
}

func PreBusiness(successList *[]structs.Area) string {
	cityJs := "https://j.i8tq.com/weather2020/search/city.js"
	citys := client.FetchString(cityJs)
	areaList := convertData.StructConvert(citys)
	return DoBusiness(successList, areaList)
}
