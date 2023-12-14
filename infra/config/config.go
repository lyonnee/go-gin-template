package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var instance Config

func Load(env string) {
	// 使用viper作为配置加载中间件
	workDir, _ := os.Getwd()
	viper.SetConfigName("config." + env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("Config file not found; ignore error if desired")
		} else {
			log.Fatal("Config file was found but another error was produced")
		}
	}

	viper.Unmarshal(&instance)
}

func App() AppConfig {
	return instance.App
}

func Log() LogConfig {
	return instance.Log
}

func Mysql() MysqlConfig {
	return instance.Mysql
}
