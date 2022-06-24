package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func initViper(env string) {
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
}
