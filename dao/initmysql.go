package dao

import (
	"fmt"
	"ginblog1/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB
var err error


func InitDb() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil{
		fmt.Printf("连接数据库失败，请检查参数: %s", err)
		return db, err
	}
	//db.Close()
	sqlDb, _ := db.DB()
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetConnMaxIdleTime(10 * time.Minute)
	sqlDb.SetConnMaxLifetime(time.Hour)
	return db, nil
}
