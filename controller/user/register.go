package user

import (
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
// @Summary     注册
// @Description 输入账户密码用户信息注册
// @Tags        user
// @Accept      json
// @Produce     json
// @Param       RegisterIM body RegisterIM true "register input model"
// @Success     200     {object} response.Response[string]
// @Router      /user/user [post]
func (c *UserController) registerHandler(ctx *gin.Context) {
	var input RegisterIM
	if err := ctx.ShouldBind(&input); err != nil {
		response.Fail(ctx, response.CODE_INVALID_BODY_ARGUMENT, "")
		return
	}

	uid, err := c.UserLogic.Register(ctx.Request.Context(), input.Name, input.Age, input.PhoneNumber, input.Password)
	if err != nil {
		response.Fail(ctx, response.CODE_PROCESSING_REQUEST_FAILURE, "")
		return
	}

	response.Ok(ctx, gin.H{
		"uid": uid,
	})
}
