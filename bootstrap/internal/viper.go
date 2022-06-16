package internal

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/soonio/pupil/app"
	"github.com/spf13/viper"
)

func Viper(config string) {
	v := viper.New()
	v.SetConfigFile(config)
	err := v.ReadInConfig()
	if err == nil {
		v.WatchConfig()
		v.OnConfigChange(func(e fsnotify.Event) {
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
