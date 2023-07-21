package main

import (
	"TeamToDo/global"
	"TeamToDo/global/logger"
	"TeamToDo/initialize"
	"TeamToDo/routes"
)

func main() {
	global.Vp = initialize.Viper()
	global.Logger = initialize.Zap()
	global.Sql = initialize.ConnectMysql()
	logger.Info("server run success on")

	// 创建路由引擎
	r := routes.SetupRoutes()
	// 开始运行
	go initialize.Scheduler()
	_ = r.Run(":8080")
}
