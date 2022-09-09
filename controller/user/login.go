package user

import (
	"github.com/LyonNee/app-layout/pkg/log"
	"github.com/LyonNee/app-layout/pkg/response"
	"github.com/LyonNee/app-layout/pkg/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LoginIM struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type LoginOM struct {
	Uid   uint64 `json:"uid"`
	Token string `json:"token"`
}

// Login
// @Summary     登录
// @Description 通过账户密码登录
// @Tags        user
// @Accept      json
// @Produce     json
// @Param       LoginIM body     LoginIM true "login input model"
// @Success     200     {object} response.Response[LoginOM]
// @Router      /user/session [post]
func (c *UserController) loginHandler(ctx *gin.Context) {
	var input LoginIM
	if err := ctx.ShouldBind(&input); err != nil {
		response.Fail(ctx, response.CODE_INVALID_BODY_ARGUMENT, "")
		return
	}

	uid, name, phoneNum, err := c.UserLogic.Login(ctx.Request.Context(), input.PhoneNumber, input.Password)
	if err != nil {
		response.Fail(ctx, response.CODE_PROCESSING_REQUEST_FAILURE, "")
		return
	}

	token, err := util.GenToken(uid, name, phoneNum)
	if err != nil {
		log.ZapLogger().Error("生成JWT失败", zap.Error(err))
		response.Ok(ctx, "")
		return
	}

	response.Ok(ctx, LoginOM{
		Uid:   uid,
		Token: token,
	})
}
