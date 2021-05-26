package logger

import "github.com/rs/zerolog"

// Config defines the config for logger middleware
type Config struct {
	Logger *zerolog.Logger
	// UTC a boolean stating whether to use UTC time zone or local.
	UTC         bool
	Env         string
	LogFilePath string
	LogFileName string
}
