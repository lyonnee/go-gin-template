package webserver

import (
	"github.com/gin-gonic/gin"
	"github.com/lyonnee/go-gin-template/docs"
	"github.com/lyonnee/go-gin-template/infra/log"
	"github.com/lyonnee/go-gin-template/webserver/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 创建并初始化router
func newRouter(env string) *gin.Engine {
	log.Debug("Init Router")

	switch env {
	case "prod":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
	r := gin.New()

	// 注册中间件
	r.Use(middleware.GinLogger(log.Logger()))
	r.Use(middleware.GinRecovery(log.Logger(), true))
	r.Use(middleware.Cros())

	// swag 初始化
	docs.SwaggerInfo.Title = "app-layout Web API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.NoRoute(func(c *gin.Context) {
		c.String(404, "未找到路由地址")
	})

	r.HandleMethodNotAllowed = true
	r.NoMethod(func(c *gin.Context) {
		c.String(404, "错误调用方法")
	})

	return r
}
