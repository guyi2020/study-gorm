package model

import (
	"study-gorm/global"
)

type JwtBlacklist struct {
	global.GvaModel
	Jwt string `gorm:"type:text;comment:jwt"`
}
