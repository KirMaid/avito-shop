package http

import (
	"avitoshop/internal/app/usecases"
	buymerch "avitoshop/internal/app/usecases/buy_merch"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BuyMerchHandler struct {
	buyMerchUseCase buymerch.BuyMerchUseCase
}

func NewBuyMerchHandler(buyMerchUseCase buymerch.BuyMerchUseCase) *BuyMerchHandler {
	return &BuyMerchHandler{buyMerchUseCase: buyMerchUseCase}
}

func (bmh *BuyMerchHandler) BuyMerch(c *gin.Context) {
	usernameInterface, exists := c.Get("username")
	if exists == false {
		c.JSON(http.StatusUnauthorized, gin.H{StatusError: ErrUsernameNotFoundInContext})
		return
	}

	username, ok := usernameInterface.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{StatusError: ErrUsernameInvalidFormat})
		return
	}

	merchName := c.Param("item")
	if merchName == "" {
		c.JSON(http.StatusBadRequest, gin.H{StatusError: "item parameter is required"})
		return
	}

	err := bmh.buyMerchUseCase.BuyMerch(c.Request.Context(), username, merchName)
	if err != nil {
		if errors.Is(err, usecases.ErrInsufficientFunds) {
			c.JSON(http.StatusBadRequest, gin.H{StatusError: err.Error()})
		}
		c.JSON(http.StatusInternalServerError, gin.H{StatusError: err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
