package log

import (
	"sync"

	"github.com/lyonnee/go-gin-template/infra/config"
	"go.uber.org/zap"
)

func Initialize() {
	once.Do(
		func() {
			zapLogger = newZap(
				config.Log().Filename,
				config.Log().Level,
				config.Log().MaxSize,
				config.Log().MaxBackups,
				config.Log().MaxAge,
			)
		},
	)
}

func Sync() {
	syncZap()
}

var once sync.Once
var zapLogger *zap.Logger

func Logger() *zap.Logger {
	if zapLogger == nil {
		once.Do(
			func() {
				zapLogger = newDefaultZap()
			},
		)
	}
	return zapLogger
}

func Debug(msg string, fields ...zap.Field) {
	Logger().Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	Logger().Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	Logger().Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Logger().Error(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	Logger().DPanic(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	Logger().Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	Logger().Fatal(msg, fields...)
}
