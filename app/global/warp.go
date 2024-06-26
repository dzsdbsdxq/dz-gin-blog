package global

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type ResponseBody[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    *T     `json:"data,omitempty"`
}

func (er ResponseBody[T]) Error() string {
	return fmt.Sprintf("%d:%s", er.Code, er.Message)
}
func ErrorResponse(message string) *ResponseBody[any] {
	return &ResponseBody[any]{
		Code:    -1,
		Message: message,
	}
}
func SuccessResponse() *ResponseBody[any] {
	return &ResponseBody[any]{
		Code:    0,
		Message: "success",
	}
}
func SuccessResponseWithData[T any](data T) *ResponseBody[T] {
	return &ResponseBody[T]{
		Code:    0,
		Message: "success",
		Data:    &data,
	}
}
func NewResponseBody[T any](code int, message string, data T) *ResponseBody[T] {
	return &ResponseBody[T]{
		Code:    code,
		Message: message,
		Data:    &data,
	}
}

type HttpCodeError int

type ErrorResponseBody struct {
	HttpCode int
	Message  string
}

func NewErrorResponseBody(httpCode int, message string) ErrorResponseBody {
	return ErrorResponseBody{
		HttpCode: httpCode,
		Message:  message,
	}
}

func (er ErrorResponseBody) Error() string {
	return er.Message
}

func ErrorHandler(ctx *gin.Context, err error) {
	l := slog.Default().With("X-Request-ID", ctx.GetString("X-Request-ID"))
	l.ErrorContext(ctx, err.Error())
	G_DZ_LOG.Errorf("stack trace: \n%+v\n", err)
	ctx.JSON(http.StatusOK, ErrorResponse(err.Error()))
}
func Wrap[T any](fn func(ctx *gin.Context) (T, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result, err := fn(ctx)
		if err != nil {
			ErrorHandler(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, result)
	}
}
func WrapWithBody[T any, R any](fn func(ctx *gin.Context, req R) (T, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req R
		bodyErr := ctx.Bind(&req)
		if bodyErr != nil {
			ErrorHandler(ctx, bodyErr)
			return
		}
		result, err := fn(ctx, req)
		if err != nil {
			ErrorHandler(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, result)
	}
}
