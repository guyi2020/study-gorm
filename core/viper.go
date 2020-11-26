package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"study-gorm/global"
	"study-gorm/utils"
)

func Viper(path ...string) *viper.Viper{
	config := utils.ConfigFile
	v := viper.New()
	v.SetConfigFile(config)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.GvaConfig); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.GvaConfig); err != nil {
		fmt.Println(err)
	}
	return v
}