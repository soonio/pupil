package internal

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/soonio/pupil/app"
	"github.com/spf13/viper"
)

// Viper 初始化配置解析工具
func Viper(config string) {
	v := viper.New()
	v.SetConfigFile(config)
	err := v.ReadInConfig()
	if err == nil {
		v.WatchConfig()
		v.OnConfigChange(func(e fsnotify.Event) {
			// TIPS: 日志会打印到控制台(或者进程监控的日志中)
			if err = v.Unmarshal(&app.Config); err != nil {
				fmt.Println("config file format error: " + err.Error())
			} else {
				fmt.Println("reload config success.")
			}
		})
		err = v.Unmarshal(&app.Config)
	}
	if err != nil {
		panic(err)
	}
}
