package dao

import (
	"server/config"
	"time"

	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
)

func init() {
	Db, err = gorm.Open("mysql", config.MysqlDb)
	if err != nil {
		// logger.Write("error")
		return
	}
	if Db.Error != nil {
		return
	}
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOPenConns(1000)
	Db.Db().SetConnMaxLifetime(time.Hour)

}
