package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lyonnee/go-gin-template/api/controller"
)

type IController interface {
	RegisterRoute(r *gin.Engine)
}

var controllers = make(map[string]IController)

// 初始化controller
func Register(r *gin.Engine) {
	controllers["user_ctrl"] = controller.NewUserCtrl()

	for _, v := range controllers {
		v.RegisterRoute(r)
	}
}
