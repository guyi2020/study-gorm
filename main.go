package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	Id  int
	Num int
}

func main() {
	// 初始化配置
	//global.GvaVp = core.Viper()
	//// 连接数据库
	//global.GvaLog = core.Zap()           // 初始化zap日志库
	//global.GvaDb = initialize.Gorm()
	//initialize.MysqlTables(global.GvaDb)
	//db, _ := global.GvaDb.DB()
	//defer db.Close()
	//
	//core.RunWindowsServer()

	dsn := "root:tifenbao2014@tcp(127.0.0.1:3306)/study_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{}); err != nil {
		fmt.Println(err)
	} else {
		db.AutoMigrate(&Product{})
	}

}
