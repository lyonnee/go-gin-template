package config

type Config struct {
	App   AppConfig
	Log   LogConfig
	Mysql MysqlConfig
}

type AppConfig struct {
	Port string
}

type LogConfig struct {
	Level      string
	Filename   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
}

type MysqlConfig struct {
	DSN string
}
