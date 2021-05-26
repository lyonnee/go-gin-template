package logger

import (
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)


// SetLogger initializes the logging middleware.
func SetGinLogger(config ...Config) gin.HandlerFunc {
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

	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}

		c.Next()

		end := time.Now()

		msg := "Request"
		if len(c.Errors) > 0 {
			msg = c.Errors.String()
		}

		dumplogger := log.Logger.With().
			Int("status", c.Writer.Status()).
			Str("method", c.Request.Method).
			Str("path", path).
			Str("ip", c.ClientIP()).
			Dur("latency", end.Sub(start)).
			Logger()

		switch {
		case c.Writer.Status() >= http.StatusBadRequest && c.Writer.Status() < http.StatusInternalServerError:
			{
				dumplogger.Warn().
					Msg(msg)
			}
		case c.Writer.Status() >= http.StatusInternalServerError:
			{
				dumplogger.Error().
					Msg(msg)
			}
		default:
			dumplogger.Info().
				Msg(msg)
		}
	}
}
