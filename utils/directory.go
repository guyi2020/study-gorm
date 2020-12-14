package utils

import (
	"go.uber.org/zap"
	"os"
	"study-gorm/global"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateDir
//@description: 批量创建文件夹
//@param: dirs ...string
//@return: err error

func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			global.GvaLog.Debug("create directory" + v)
			err = os.MkdirAll(v, os.ModePerm)
			if err != nil {
				global.GvaLog.Error("create directory"+ v, zap.Any(" error:", err))
			}
		}
	}
	return err
}
