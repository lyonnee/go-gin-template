package user

import (
	"github.com/LyonNee/app-layout/router"
)

// handler注册路由
func Initialize() {
	r := router.Get().Group("/user")

	r.POST("/session", loginHandler)
	r.POST("", registerHandler)
	// 新的handler在这里注册路由
}
