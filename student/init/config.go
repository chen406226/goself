package init

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"student/global"
)

const defaultConfigFile = "config.yaml"

func init() {
	v := viper.New()

	//v.AddConfigPath("/")
	//v.SetConfigName("config")
	//v.SetConfigType("yaml")	==👇

	v.SetConfigFile(defaultConfigFile)
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file : #{err} \n"))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:",e.Name)
		if err := v.Unmarshal(*global.GL_CONFIG); err != nil {
			fmt.Println("Error23",err)
		}
	})

	//反序列化未struct -> GL_CONFIG
	if err := v.Unmarshal(&global.GL_CONFIG); err != nil {
		fmt.Println("Error27",err)
	}
	global.GL_VP = v
	fmt.Println("config.go-init:after")
}

