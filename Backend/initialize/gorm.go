package initialize

import (
	"TeamToDo/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectMysql() *gorm.DB {
	dsn := global.Server.MySQL.Dsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		global.Logger.Errorf("mysql connect error %s\n", err.Error())
	}
	return db
}
