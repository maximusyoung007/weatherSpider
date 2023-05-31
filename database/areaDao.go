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
	dataSource := mysqlConf.Username + ":" + mysqlConf.Password + "@tcp(127.0.0.1:3306)/" + mysqlConf.Database + "?charset=utf8&parseTime=True"
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
		sqlStr := "insert into airCondition(areaId, name, airCondition, systemTime) values (?,?,?,?) "
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
