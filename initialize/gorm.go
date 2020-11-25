package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	return GormMysql()
}

func GormMysql() *gorm.DB  {
	//mysqlConfig := mysql.Config{
	//
	//}
	dsn := "root:tifenbao2014@tcp(127.0.0.1:3306)/study_gorm?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	return db
}
