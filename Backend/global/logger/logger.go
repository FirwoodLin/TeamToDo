package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func init() {
	var logger *zap.Logger
	cfg := zap.NewDevelopmentConfig()

	// 设置不同级别的颜色
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	logger, _ = cfg.Build()

	Logger = logger.Sugar()
}

func Info(args ...interface{}) {
	Logger.Info(args...)
}
