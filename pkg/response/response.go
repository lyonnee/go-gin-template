package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code uint16      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// base response
func NewResponse(c *gin.Context, code uint16, msg string, data interface{}) {
	c.JSON(
		http.StatusOK,
		Response{
			Code: code,
			Msg:  msg,
			Data: data,
		},
	)
}

func Ok(c *gin.Context, data interface{}) {
	NewResponse(c, CODE_CALL_SUCCESS, "", data)
}

func Fail(c *gin.Context, code uint16, msg string) {
	NewResponse(c, code, msg, nil)
}

// extensions response
func InvalidQueryArgument(c *gin.Context) {
	NewResponse(c, CODE_INVALID_QUERY_ARGUMENT, "invalid query argument", nil)
}

func InvalidPathArgument(c *gin.Context) {
	NewResponse(c, CODE_INVALID_PATH_ARGUMENT, "invalid path argument", nil)
}

func InvalidBodyArgument(c *gin.Context) {
	NewResponse(c, CODE_INVALID_BODY_ARGUMENT, "invalid body argument", nil)
}
