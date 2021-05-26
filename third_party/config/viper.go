package config

import (
	"github.com/spf13/viper"
	"os"
)

func Load(filename string,env string) {
	workDir, _ := os.Getwd()

	switch env {
	case "test":
		viper.SetConfigName(filename + ".test")
		break
	case "dev":
		viper.SetConfigName(filename + ".dev")
		break
	default:
		viper.SetConfigName(filename)
	}

	viper.SetConfigType("yml")               // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(workDir + "/configs") // optionally look for config in the working directory

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			panic(err)
		} else {
			// Config file was found but another error was produced
		}
	}
}
