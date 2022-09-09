package controller

import (
	"github.com/LyonNee/app-layout/controller/user"
	"github.com/LyonNee/app-layout/logic"
	"github.com/gin-gonic/gin"
)

// 初始化controller
func Initialize(r *gin.Engine) {
	userController := user.UserController{
		UserLogic: &logic.UserLogic{},
	}
	userController.RegisterRoute(r)
}
