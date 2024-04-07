package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Respuesta exitosa
func Success(ctx *gin.Context, status int, data interface{}, message string) {
	ctx.JSON(status, response{
		Message: message,
		Data:    data,
	})
}

// Respuesta con error
func Failure(ctx *gin.Context, status int, err error) {
	ctx.JSON(status, errorResponse{
		Message: err.Error(),
		Status:  status,
		Code:    http.StatusText(status),
	})
}
