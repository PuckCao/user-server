package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbCon *gorm.DB

//初始化mysql链接
func InitMysql() {
	//gorm初始化数据库链接
	dsn := "root:1997@tcp(127.0.0.1:3306)/mall?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	dbCon = db
}
