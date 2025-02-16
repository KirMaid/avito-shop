package http

import (
	"avitoshop/internal/app/usecases"
	sendcoins "avitoshop/internal/app/usecases/send_coins"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

var badRequestErrors = []error{
	usecases.ErrInsufficientFunds,
	sendcoins.ErrAmountMustBePositive,
	sendcoins.ErrSamePersons,
	sendcoins.ErrFailedToGetReceiver,
}

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
		Amount string `json:"amount" binding:"required,number"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			c.JSON(http.StatusBadRequest, gin.H{StatusError: "Invalid input"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{StatusError: ErrInvalidRequestBody.Error()})
		}
		return
	}

	amount, err := strconv.Atoi(request.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{StatusError: "amount must be a valid integer"})
		return
	}

	err = sch.sendCoinsUseCase.SendCoins(c.Request.Context(), senderUsername, request.ToUser, amount)
	if err != nil {
		for _, targetErr := range badRequestErrors {
			if errors.Is(err, targetErr) {
				c.JSON(http.StatusBadRequest, gin.H{StatusError: err.Error()})
				return
			}
		}
		c.JSON(http.StatusInternalServerError, gin.H{StatusError: err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
