package ginserver

import (
	"context"
	"net/http"

	"github.com/LyonNee/app-layout/controller"
	"github.com/LyonNee/app-layout/docs"
	"github.com/LyonNee/app-layout/pkg/log"
	"github.com/LyonNee/app-layout/router"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"go.uber.org/zap"
)

var srv *http.Server

func Run() {
	log.ZapLogger().Debug("Web API Service run")
	// 获取路由
	router := router.Get()

	// 初始化handler(注册路由)
	controller.Initialize()

	// swag 初始化
	docs.SwaggerInfo.Title = "app-layout Web API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	srv = &http.Server{
		Addr:    viper.GetString("app.port"),
		Handler: router,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.ZapLogger().Fatal("Web API Service Not Run", zap.Error(err))
	}
}

func Shutdown(ctx context.Context) {
	srv.Shutdown(ctx)
}
