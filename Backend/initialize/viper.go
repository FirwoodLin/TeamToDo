package initialize

import (
	"TeamToDo/global"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper() *viper.Viper {
	var config string
	config = "./global/config.yaml"
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error global file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("global file changed:", e.Name)
		if err := v.Unmarshal(&global.Server); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.Server); err != nil {
		fmt.Println(err)
	}
	return v
}
