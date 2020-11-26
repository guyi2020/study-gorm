package main

import (
	"fmt"
	"study-gorm/core"
	"study-gorm/global"
	"study-gorm/initialize"
)

func main()  {
	// 初始化配置
	global.GvaVp = core.Viper()
	// 连接数据库
	global.GvaDb = initialize.Gorm()
	db, _ := global.GvaDb.DB()
	defer db.Close()
	fmt.Println(global.GvaDb)

	Router := initialize.Routers()
	fmt.Println(Router)
}


