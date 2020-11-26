package global

import (
	"github.com/spf13/viper"
	"study-gorm/config"

	"gorm.io/gorm"
)

var (
	GvaDb     *gorm.DB
	GvaConfig config.Server
	GvaVp     *viper.Viper
)
