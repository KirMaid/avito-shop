package http

import (
	"avitoshop/internal/app/usecases"
	usecase "avitoshop/internal/app/usecases/user_info"
	"avitoshop/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter(
	handler *gin.Engine,
	l logger.Interface,
	authMiddleware gin.HandlerFunc,
	auth usecases.AuthUseCase,
	userInfo usecase.UserInfoUseCase,
) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// K8s probe
	handler.GET("/health", func(c *gin.Context) { c.Status(http.StatusOK) })

	//TODO Внедрить Prometheus

	// Prometheus metrics
	//handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Routers
	authHandler := NewAuthHandler(auth)
	userInfoHandler := NewUserInfoHandler(userInfo)
	handler.POST("api//auth", authHandler.Auth)
	handler.GET("/api/info", authMiddleware, userInfoHandler.GetInfo)
	//TODO Дописать
	handler.POST("/api/sendCoin", authMiddleware, nil)
	handler.GET("/api/buy/:item", authMiddleware, nil)
}
