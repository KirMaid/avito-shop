package http

import (
	"avitoshop/internal/app/usecases"
	sendcoins "avitoshop/internal/app/usecases/send_coins"
	"errors"
	"github.com/gin-gonic/gin"
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
		c.JSON(http.StatusUnauthorized, gin.H{StatusError: ErrUsernameNotFoundInContext})
		return
	}

	senderUsername, ok := senderUsernameInterface.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{StatusError: ErrUsernameInvalidFormat})
		return
	}

	var request struct {
		ToUser string `json:"toUser" binding:"required"`
		Amount string `json:"amount" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{StatusError: ErrInvalidRequestBody})
		return
	}

	amount, err := strconv.Atoi(request.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{StatusError: "amount must be a valid integer"})
		return
	}

	err = sch.sendCoinsUseCase.SendCoins(c.Request.Context(), senderUsername, request.ToUser, amount)
	if err != nil {
		if errors.Is(err, usecases.ErrInsufficientFunds) {
			c.JSON(http.StatusBadRequest, gin.H{StatusError: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{StatusError: err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
