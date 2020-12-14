package service

import (
	"study-gorm/global"
	"study-gorm/model"
)

func FindUserById(id int) (err error, user *model.SysUser) {
	var u model.SysUser
	err = global.GvaDb.Where("`id` = ?", id).First(&u).Error
	return err, &u
}
