package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lyonnee/go-gin-template/api/iomodel"
	"github.com/lyonnee/go-gin-template/api/response"
	"github.com/lyonnee/go-gin-template/infra/log"
	"github.com/lyonnee/go-gin-template/logic"
	"github.com/lyonnee/go-gin-template/util"
	"go.uber.org/zap"
)

type UserCtrl struct {
	userLogic *logic.UserLogic
}

func NewUserCtrl() *UserCtrl {
	return &UserCtrl{
		userLogic: logic.NewUserLogic(),
	}
}

func (c *UserCtrl) RegisterRoute(r *gin.Engine) {
	group := r.Group("/api/user")

	group.POST("/session", c.loginHandler)
	group.POST("/user", c.registerHandler)
}

// Login
//
//	@Summary		注册
//	@Description	输入账户密码用户信息注册
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			RegisterIM	body		iomodel.RegisterIM	true	"register input model"
//	@Success		200			{object}	response.Response[string]
//	@Router			/user/user [post]
func (c *UserCtrl) registerHandler(ctx *gin.Context) {
	var input iomodel.RegisterIM
	if err := ctx.ShouldBind(&input); err != nil {
		response.Fail(ctx, response.CODE_INVALID_BODY_ARGUMENT, "")
		return
	}

	uid, err := c.userLogic.Register(ctx.Request.Context(), input.Name, input.Age, input.PhoneNumber, input.Password)
	if err != nil {
		response.Fail(ctx, response.CODE_PROCESSING_REQUEST_FAILURE, "")
		return
	}

	response.Ok(ctx, gin.H{
		"uid": uid,
	})
}

// Login
//
//	@Summary		登录
//	@Description	通过账户密码登录
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			LoginIM	body		iomodel.LoginIM	true	"login input model"
//	@Success		200		{object}	response.Response[iomodel.LoginOM]
//	@Router			/user/session [post]
func (c *UserCtrl) loginHandler(ctx *gin.Context) {
	var input iomodel.LoginIM
	if err := ctx.ShouldBind(&input); err != nil {
		response.Fail(ctx, response.CODE_INVALID_BODY_ARGUMENT, "")
		return
	}

	uid, name, phoneNum, err := c.userLogic.Login(ctx.Request.Context(), input.PhoneNumber, input.Password)
	if err != nil {
		response.Fail(ctx, response.CODE_PROCESSING_REQUEST_FAILURE, "")
		return
	}

	token, err := util.GenToken(uid, name, phoneNum)
	if err != nil {
		log.Error("生成JWT失败", zap.Error(err))
		response.Ok(ctx, "")
		return
	}

	response.Ok(ctx, iomodel.LoginOM{
		Uid:   uid,
		Token: token,
	})
}
