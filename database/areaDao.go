package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"weatherSpider/structs"
)

var db *sql.DB

func initDB() (err error) {
	db, err = sql.Open("mysql", "root:jnhfj@2009@tcp(127.0.0.1:3306)/learnsomething?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Println("数据库链接错误", err)
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	//延迟到函数结束关闭链接
	//defer db.Close()
	return nil
}

func InsertRow(areaList []structs.Area) {
	initDB()
	for i := 0; i < len(areaList); i++ {
		area := areaList[i]
		sqlStr := "insert into airCondition(areaId, name, airCondition, systemTime) values (?,?,?,?) "
		ret, err := db.Exec(sqlStr, area.AreaId, area.NameCN, area.AirCondition, time.Now())
		if err != nil {
			fmt.Printf("insert failed, err:%v\n", err)
			return
		}
		theID, err := ret.LastInsertId() // 新插入数据的id
		if err != nil {
			fmt.Printf("get lastinsert ID failed, err:%v\n", err)
			return
		}
		fmt.Printf("insert success, the id is %d.\n", theID)
	}
}
