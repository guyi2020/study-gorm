package main

import (
	"fmt"
	"study-gorm/global"
	"study-gorm/initialize"
)

func main()  {
	global.GVA_DB = initialize.Gorm()
	fmt.Println(global.GVA_DB)
}


