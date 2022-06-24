package router

import (
	"sync"

	"github.com/LyonNee/app-layout/pkg/log"
	"github.com/LyonNee/app-layout/pkg/middleware"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
	once   sync.Once
)

func newRouter() {
	log.ZapLogger().Debug("Init Router")
	router = gin.New()

	router.Use(middleware.GinLogger(log.ZapLogger()))
	router.Use(middleware.GinRecovery(log.ZapLogger(), true))
	router.Use(middleware.Cros())

	router.NoRoute(func(c *gin.Context) {
		c.String(404, "未找到路由地址")
	})

	router.HandleMethodNotAllowed = true
	router.NoMethod(func(c *gin.Context) {
		c.String(404, "错误调用方法")
	})
}

func Get() *gin.Engine {
	if router == nil {
		once.Do(
			newRouter,
		)
	}
	return router
}

func GetV1() *gin.RouterGroup {
	r := Get()
	v1 := r.Group("/api/v1")

	return v1
}
