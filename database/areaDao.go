package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"time"
	"weatherSpider/conf"
	"weatherSpider/logu"
	"weatherSpider/structs"
)

var db *sql.DB
var logger = &logu.Logger
var mysqlConf = &conf.Conf.Mysql

func initDB() (err error) {
	dataSource := mysqlConf.Username + ":" + mysqlConf.Password + "@tcp(" + mysqlConf.Ip + ":3306)/" + mysqlConf.Database + "?parseTime=true&loc=Local"
	db, err = sql.Open("mysql", dataSource)
	if err != nil {
		(*logger).WithFields(logrus.Fields{"error": err}).Error("数据库连接错误")
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
		sqlStr := "insert into weather(areaId, name, airCondition, systemTime) values (?,?,?,?) "
		ret, err := db.Exec(sqlStr, area.AreaId, area.NameCN, area.AirCondition, time.Now())
		if err != nil {
			(*logger).WithFields(logrus.Fields{"error": err}).Error("insert failed")
			return
		}
		theID, err := ret.LastInsertId() // 新插入数据的id
		if err != nil {
			(*logger).WithFields(logrus.Fields{
				"error": err,
			}).Error("get lastinsert ID failed")
			return
		}
		(*logger).WithFields(logrus.Fields{"id": theID}).Info("insert success")
	}
}

func GetRow() []structs.Area {
	initDB()
	areaList := make([]structs.Area, 0)
	sqlStr := "select name, areaId, ROUND(avg(airCondition), 2) avg from weather group by name order by avg(airCondition)"
	rows, err := db.Query(sqlStr)
	if err != nil {
		(*logger).Error("selected failed")
		return []structs.Area{}
	}
	defer rows.Close()
	for rows.Next() {
		var a structs.Area
		err := rows.Scan(&a.NameCN, &a.AreaId, &a.AvgAir)
		if err != nil {
			(*logger).Error("scan failed")
			return []structs.Area{}
		}
		areaList = append(areaList, a)
	}
	return areaList
}
