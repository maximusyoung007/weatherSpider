package convertData

import (
	"encoding/json"
	"fmt"
	"weatherSpider/structs"
)

func StructConvert(s string) []structs.Area {
	areaList1 := make([]structs.Area, 0)

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
	return areaList1
}
