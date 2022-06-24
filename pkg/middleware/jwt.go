package middleware

import (
	"strings"

	"github.com/LyonNee/app-layout/pkg/response"
	"github.com/LyonNee/app-layout/pkg/util"

	"github.com/gin-gonic/gin"
)

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.Fail(ctx, response.CODE_NOT_TOKEN, "无权限访问，请求未携带token")
			ctx.Abort() //结束后续操作
			return
		}

		//按空格拆分
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.Fail(ctx, response.CODE_TOKEN_FORMAT_INCORRECT, "请求头中auth格式有误")
			ctx.Abort()
			return
		}

		//解析token包含的信息
		claims, err := util.ParseToken(parts[1])
		if err != nil {
			response.Fail(ctx, response.CODE_TOKEN_INVALID, "无效的Json Web Token")
			ctx.Abort()
			return
		}

		// 将当前请求的claims信息保存到请求的上下文c上
		ctx.Set("claims", claims)
		ctx.Next() // 后续的处理函数可以用过ctx.Get("claims")来获取当前请求的用户信息

	}
}
