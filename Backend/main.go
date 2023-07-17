package main

import (
	"TeamToDo/global"
	"TeamToDo/initialize"
)

func main() {
	global.Vp = initialize.Viper()
	global.Logger = initialize.Zap()
	global.Sql = initialize.ConnectMysql()
	//fmt.Printf("%v\n", global.Server)
	global.Logger.Info("server run success on")
	//global.Logger.Debug("server run success on ")

}
