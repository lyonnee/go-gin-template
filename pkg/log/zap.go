package log

import (
	"os"

	"github.com/LyonNee/app-layout/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zapLogger *zap.Logger

func ZapLogger() *zap.Logger {
	if zapLogger == nil {
		zapLogger, _ = zap.NewProduction()
	}
	return zapLogger
}

func initZap() {
	writeSyncer := getLogWriter(
		config.Instance().Log.Filename,
		config.Instance().Log.MaxSize,
		config.Instance().Log.MaxBackups,
		config.Instance().Log.MaxAge,
	)

	encoder := getEncoder()
	var level = new(zapcore.Level)
	err := level.UnmarshalText([]byte(config.Instance().Log.Level))

	if err != nil {
		return
	}

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level),
		zapcore.NewCore(encoder, writeSyncer, level),
	)

	zapLogger = zap.New(core, zap.AddCaller())
}

func syncZap() {
	zapLogger.Sync()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}

	return zapcore.AddSync(lumberJackLogger)
}
