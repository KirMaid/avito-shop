package http

import (
	sendcoins "avitoshop/internal/app/usecases/send_coins"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type SendCoinsHandler struct {
	sendCoinsUseCase sendcoins.SendCoinsUseCase
}

func NewSendCoinsHandler(sendCoinsUseCase sendcoins.SendCoinsUseCase) *SendCoinsHandler {
	return &SendCoinsHandler{sendCoinsUseCase: sendCoinsUseCase}
}

func (sch *SendCoinsHandler) SendCoins(c *gin.Context) {
	senderUsernameInterface, exists := c.Get("username")
	if exists == false {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": "senderUsername not found in context"})
		return
	}

	senderUsername, ok := senderUsernameInterface.(string)
	if !ok {
		log.Printf("Гойда ошибка Нулевая")
		c.JSON(http.StatusInternalServerError, gin.H{"errors": "senderUsername is not an string"})
		return
	}

	var request struct {
		ToUser string `json:"toUser" binding:"required"`
		Amount string `json:"amount" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "invalid request body"})
		return
	}

	amount, err := strconv.Atoi(request.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "amount must be a valid integer"})
		return
	}

	err = sch.sendCoinsUseCase.SendCoins(c.Request.Context(), senderUsername, request.ToUser, amount)
	if err != nil {
		log.Printf("Гойда ошибка Первая")
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
