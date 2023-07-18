package initialize

import (
	"TeamToDo/global"
	"TeamToDo/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ConnectMysql 建立 Mysql 连接
func ConnectMysql() *gorm.DB {
	dsn := global.Server.MySQL.Dsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		global.Logger.Errorf("mysql connect error %s\n", err.Error())
	}
	initTables(db)
	return db
}

// initTables 自动迁移表
func initTables(db *gorm.DB) {
	tables := []interface{}{
		&model.User{},
		&model.EmailVerification{},
	}
	err := db.AutoMigrate(tables...)
	if err != nil {
		global.Logger.Errorf("register table failed: %v\n", err)

	}
}
