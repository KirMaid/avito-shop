package http

import (
	auth "avitoshop/internal/app/usecases/auth"
	sendcoin "avitoshop/internal/app/usecases/send_coins"
	userinfo "avitoshop/internal/app/usecases/user_info"
	"avitoshop/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

const StatusError = "errors"

func NewRouter(
	handler *gin.Engine,
	l logger.Interface,
	authMiddleware gin.HandlerFunc,
	auth auth.AuthUseCase,
	userInfo userinfo.UserInfoUseCase,
	sendCoins sendcoin.SendCoinsUseCase,
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
	sendCoinsHandler := NewSendCoinsHandler(sendCoins)

	handler.POST("api/auth", authHandler.Auth)
	handler.GET("/api/info", authMiddleware, userInfoHandler.GetInfo)
	//TODO Дописать
	handler.POST("/api/sendCoin", authMiddleware, sendCoinsHandler.SendCoins)
	handler.GET("/api/buy/:item", authMiddleware, nil)
}
