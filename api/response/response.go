package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// base response
type Response[T any] struct {
	Code uint16 `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

func New[T any](code uint16, msg string, data T) *Response[T] {
	return &Response[T]{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func Ok[T any](c *gin.Context, data T) {
	resp := New(CODE_CALL_SUCCESS, "", data)
	c.JSON(
		http.StatusOK,
		resp,
	)
}

func Fail(c *gin.Context, code uint16, msg string) {
	resp := New(code, msg, "")
	c.JSON(
		http.StatusOK,
		resp,
	)
}

// extensions response
func InvalidQueryArgument(c *gin.Context) {
	Fail(c, CODE_INVALID_QUERY_ARGUMENT, MSG_INVALID_QUERY_ARGUMENT)
}

func InvalidPathArgument(c *gin.Context) {
	Fail(c, CODE_INVALID_PATH_ARGUMENT, MSG_INVALID_PATH_ARGUMENT)
}

func InvalidBodyArgument(c *gin.Context) {
	Fail(c, CODE_INVALID_BODY_ARGUMENT, MSG_INVALID_BODY_ARGUMENT)
}
