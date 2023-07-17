package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Vp     *viper.Viper
	Logger *zap.SugaredLogger
	Server Config
	Sql    *gorm.DB
)
