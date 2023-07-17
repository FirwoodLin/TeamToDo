package initialize

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Zap() (sugarLogger *zap.SugaredLogger) {
	var logger *zap.Logger
	cfg := zap.NewDevelopmentConfig()

	// 设置不同级别的颜色
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	logger, _ = cfg.Build()

	sugarLogger = logger.Sugar()
	return sugarLogger
}
