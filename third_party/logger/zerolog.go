package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path"
	"path/filepath"
	"strconv"
)

func SetLogger(config ...Config){
	var newConfig Config
	if len(config) > 0 {
		newConfig = config[0]
	}

	basicLog := zerolog.ConsoleWriter{
		Out:     os.Stderr,
		NoColor: true,
		TimeFormat: "2006-01-02 15:04:05",
	}

	switch newConfig.Env {
	case "dev":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	case "test":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	default:
		// 路径脱敏
		zerolog.CallerMarshalFunc = func(file string, line int) string {
			return filepath.Base(file) + ":" + strconv.Itoa(line)
		}
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	// 日志文件
	fileName := path.Join(newConfig.LogFilePath,newConfig.LogFileName)
	basicLog.Out = &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    100, // megabytes
		MaxBackups: 10,
		MaxAge:     30, //days
		LocalTime:  true,
		Compress:   true, // disabled by default
	}

	log.Logger = zerolog.New(basicLog).With().Timestamp().Caller().Logger()
}