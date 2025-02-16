package http

import (
	"avitoshop/internal/app/usecases"
	buygood "avitoshop/internal/app/usecases/buy_good"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BuyGoodHandler struct {
	buyGoodUseCase buygood.BuyGoodUseCase
}

func NewBuyGoodHandler(buyGoodUseCase buygood.BuyGoodUseCase) *BuyGoodHandler {
	return &BuyGoodHandler{buyGoodUseCase: buyGoodUseCase}
}

func (bmh *BuyGoodHandler) BuyGood(c *gin.Context) {
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

	goodName := c.Param("item")
	if goodName == "" {
		c.JSON(http.StatusBadRequest, gin.H{StatusError: "item parameter is required"})
		return
	}

	err := bmh.buyGoodUseCase.BuyGood(c.Request.Context(), username, goodName)
	if err != nil {
		if errors.Is(err, usecases.ErrInsufficientFunds) {
			c.JSON(http.StatusBadRequest, gin.H{StatusError: err.Error()})
			return
		}
		if errors.Is(err, buygood.ErrFailedToGetGood) {
			c.JSON(http.StatusNotFound, gin.H{StatusError: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{StatusError: err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
