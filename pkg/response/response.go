package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response[T any] struct {
	Code uint16 `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

func New[T any](code uint16, msg string, data T) Response[T] {
	return Response[T]{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

// base response
func NewResponse[T any](c *gin.Context, code uint16, msg string, data T) {
	c.JSON(
		http.StatusOK,
		Response[T]{
			Code: code,
			Msg:  msg,
			Data: data,
		},
	)
}

func Ok[T any](c *gin.Context, data T) {
	NewResponse(c, CODE_CALL_SUCCESS, "", data)
}

func Fail(c *gin.Context, code uint16, msg string) {
	NewResponse(c, code, msg, "")
}

// extensions response
func InvalidQueryArgument(c *gin.Context) {
	NewResponse(c, CODE_INVALID_QUERY_ARGUMENT, "invalid query argument", "")
}

func InvalidPathArgument(c *gin.Context) {
	NewResponse(c, CODE_INVALID_PATH_ARGUMENT, "invalid path argument", "")
}

func InvalidBodyArgument(c *gin.Context) {
	NewResponse(c, CODE_INVALID_BODY_ARGUMENT, "invalid body argument", "")
}
