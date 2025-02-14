package http

import (
	"avitoshop/internal/app/entities"
	"avitoshop/internal/app/usecases"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

const StatusError = "errors"

var ErrInvalidAccessToken = errors.New("invalid auth token")
var ErrUserDoesNotExist = errors.New("user does not exist")

type Response struct {
	Status string `json:"status"`
	Msg    string `json:"message,omitempty"`
}

type AuthHandler struct {
	authUseCase usecases.AuthUseCase
}

func NewAuthHandler(authUseCase usecases.AuthUseCase) *AuthHandler {
	return &AuthHandler{authUseCase: authUseCase}
}

func (ah *AuthHandler) Auth(c *gin.Context) {
	var auth entities.Auth
	if err := c.ShouldBindJSON(&auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{StatusError: "invalid request body"})
		return
	}

	token, err := ah.authUseCase.Auth(c.Request.Context(), &auth)
	if err != nil {
		if errors.Is(err, ErrInvalidAccessToken) || errors.Is(err, ErrUserDoesNotExist) {
			c.JSON(http.StatusBadRequest, gin.H{StatusError: err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{StatusError: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
