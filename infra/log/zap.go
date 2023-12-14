package log

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newZap(filename, levelStr string, maxSize, maxBackups, maxAge int) *zap.Logger {
	writeSyncer := getLogWriter(
		filename,
		maxSize,
		maxBackups,
		maxAge,
	)

	encoder := getEncoder()
	var level = new(zapcore.Level)
	err := level.UnmarshalText([]byte(levelStr))

	if err != nil {
		return newDefaultZap()
	}

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level),
		zapcore.NewCore(encoder, writeSyncer, level),
	)

	return zap.New(core, zap.AddCaller())
}

func newDefaultZap() *zap.Logger {
	zapLogger, _ := zap.NewProduction()
	return zapLogger
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
