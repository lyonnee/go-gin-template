package config

import (
	pkgConfig "github.com/LyonNee/app-layout/pkg/config"
	"github.com/spf13/viper"
)

var instance Config

func Initialize(env string) {
	// 使用viper作为配置加载中间件
	pkgConfig.InitViper(env)
	viper.Unmarshal(&instance)
}

func GetConfig() *Config {
	return &instance
}
