package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"study-gorm/config"

	"gorm.io/gorm"
)

var (
	GvaDb     *gorm.DB
	GvaConfig config.Server
	GvaVp     *viper.Viper
	GvaLog    *zap.Logger
)
