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
	//uuid := database.NewVerifyLinkUuid("test@qq.com")
	//err := database.VerifyEmail(uuid)
	//fmt.Println(err)
	//global.Logger.Debug("server run success on ")

}
