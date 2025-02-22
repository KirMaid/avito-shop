package http

import (
	auth "avitoshop/internal/app/usecases/auth"
	buygood "avitoshop/internal/app/usecases/buy_good"
	sendcoin "avitoshop/internal/app/usecases/send_coins"
	userinfo "avitoshop/internal/app/usecases/user_info"
	"avitoshop/pkg/logger"
	"github.com/gin-gonic/gin"
)

func NewRouter(
	handler *gin.Engine,
	l logger.Interface,
	authMiddleware gin.HandlerFunc,
	auth auth.AuthUseCase,
	userInfo userinfo.UserInfoUseCase,
	sendCoins sendcoin.SendCoinsUseCase,
	buyGood buygood.BuyGoodUseCase,
) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	authHandler := NewAuthHandler(auth)
	userInfoHandler := NewUserInfoHandler(userInfo)
	sendCoinsHandler := NewSendCoinsHandler(sendCoins)
	buyGoodHandler := NewBuyGoodHandler(buyGood)

	api := handler.Group("/api")
	{
		api.POST("/auth", authHandler.Auth)
		api.GET("/info", authMiddleware, userInfoHandler.GetInfo)
		api.POST("/sendCoin", authMiddleware, sendCoinsHandler.SendCoins)
		api.GET("/buy/:item", authMiddleware, buyGoodHandler.BuyGood)
	}
}
