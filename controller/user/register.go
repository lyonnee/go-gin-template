package user

import (
	"github.com/LyonNee/app-layout/logic"
	"github.com/LyonNee/app-layout/pkg/response"
	"github.com/gin-gonic/gin"
)

type RegisterIM struct {
	Name        string `json:"name"`
	Age         uint8  `json:"age"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

// Login
// @Summary      登录
// @Description  通过账户密码登录
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        RegisterIM  body  RegisterIM  true  "register input model"
// @Router       /user/session [post]
func registerHandler(ctx *gin.Context) {
	var input RegisterIM
	if err := ctx.ShouldBind(&input); err != nil {
		response.Fail(ctx, response.CODE_INVALID_BODY_ARGUMENT, "")
		return
	}

	uid, err := logic.Register(ctx.Request.Context(), input.Name, input.Age, input.PhoneNumber, input.Password)
	if err != nil {
		response.Fail(ctx, response.CODE_PROCESSING_REQUEST_FAILURE, "")
		return
	}

	response.Ok(ctx, gin.H{
		"uid": uid,
	})
}
