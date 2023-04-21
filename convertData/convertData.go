package convertData

import (
	"encoding/json"
	"fmt"
	"time"
	"weatherSpider/structs"
)

func StructConvert(s string) []structs.Area {
	areaList1 := make([]structs.Area, 0)

	t := time.Now()
	stack := make([]string, 0)
	level4Saved := false
	level := 0
	indexI, indexJ := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '{' {
			stack = append(stack, string(s[i]))
			level++
			if level == 4 && !level4Saved {
				indexI = i
			}
		} else if s[i] == '}' {
			stack = stack[:len(stack)-1]
			if level == 4 && !level4Saved {
				indexJ = i
				city := s[indexI : indexJ+1]
				var area structs.Area
				err := json.Unmarshal([]byte(city), &area)
				areaList1 = append(areaList1, area)
				if err != nil {
					fmt.Print("城市序列化失败:", city)
				}
				level4Saved = true
			}
			level--
			if level == 2 {
				level4Saved = false
			}
		}
	}
	fmt.Println(time.Since(t))
	return areaList1

	//json转map再提取area比较慢
	//t1 := time.Now()
	//m := make(map[string]map[string]map[string]map[string]string)
	//for i := 0; i < len(s); i++ {
	//	if s[i] == '{' {
	//		s = s[i:]
	//		break
	//	}
	//}
	//err := json.Unmarshal([]byte(s), &m)
	//if err != nil {
	//	fmt.Print("json 解析成map失败", err)
	//}
	//for _, v := range m {
	//	for key, v1 := range v {
	//		for _, v2 := range v1 {
	//			if key == v2["NAMECN"] {
	//				var area structs.Area
	//				area.AreaId = v2["AREAID"]
	//				area.NameCN = v2["NAMECN"]
	//				areaList2 = append(areaList2, area)
	//			}
	//		}
	//	}
	//}
	//
	//fmt.Print(time.Since(t1))
}
