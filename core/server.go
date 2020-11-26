package core

import (
	"fmt"
	"study-gorm/initialize"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	Router := initialize.Routers()
	//fmt.Println(Router)
	address := "127.0.0.1:9999"
	s := initServer(address, Router)
	fmt.Println(s.ListenAndServe().Error())
}
