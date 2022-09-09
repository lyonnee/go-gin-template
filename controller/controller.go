package controller

import "github.com/gin-gonic/gin"

type IController interface {
	RegisterRoute(r *gin.RouterGroup)
}
