package user

import (
	"github.com/LyonNee/app-layout/logic"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserLogic *logic.UserLogic
}

func (c *UserController) RegisterRoute(r *gin.Engine) {
	group := r.Group("/api/v1/user")

	group.POST("/session", c.loginHandler)
	group.POST("/user", c.registerHandler)
}
