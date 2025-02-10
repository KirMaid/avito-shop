package delivery

import (
	"avitoshop/internal/app/auth"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, usecase auth.UseCase) {
	h := newHandler(usecase)

	router.POST("/auth", h.Auth)
}
