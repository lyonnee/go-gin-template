package app

import (
	"app-layout/third_party/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"github.com/spf13/viper"
)

func Start(env string) {

	router := gin.New()

	router.Use(cors.Default())
	router.Use(logger.SetGinLogger(logger.Config{
		LogFilePath: "logs",
		LogFileName: "webapi.log",
		Env: env,
	}))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//http://ascii.mastervb.net/text_to_ascii.php
	image :="      ___                         ___           ___                   \n     /  /\\          ___          /  /\\         /  /\\          ___     \n    /  /::\\        /__/\\        /  /::\\       /  /::\\        /__/\\    \n   /__/:/\\:\\       \\  \\:\\      /  /:/\\:\\     /  /:/\\:\\       \\  \\:\\   \n  _\\_ \\:\\ \\:\\       \\__\\:\\    /  /::\\ \\:\\   /  /::\\ \\:\\       \\__\\:\\  \n /__/\\ \\:\\ \\:\\      /  /::\\  /__/:/\\:\\_\\:\\ /__/:/\\:\\_\\:\\      /  /::\\ \n \\  \\:\\ \\:\\_\\/     /  /:/\\:\\ \\__\\/  \\:\\/:/ \\__\\/~|::\\/:/     /  /:/\\:\\\n  \\  \\:\\_\\:\\      /  /:/__\\/      \\__\\::/     |  |:|::/     /  /:/__\\/\n   \\  \\:\\/:/     /__/:/           /  /:/      |  |:|\\/     /__/:/     \n    \\  \\::/      \\__\\/           /__/:/       |__|:|~      \\__\\/      \n     \\__\\/                       \\__\\/         \\__\\|                  "

	fmt.Println(image)
	router.Run(viper.GetString("app.webapi.port"))
}