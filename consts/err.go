package consts

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义通用的错误响应结构
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// 封装统一的错误响应函数
func RespondWithError(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, ErrorResponse{
		Code:    code,
		Message: message,
	})
}
