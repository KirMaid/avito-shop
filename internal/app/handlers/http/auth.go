package http

import (
	"avitoshop/internal/app/entities"
	auth "avitoshop/internal/app/usecases/auth"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Status string `json:"status"`
	Msg    string `json:"message,omitempty"`
}

type AuthHandler struct {
	authUseCase auth.AuthUseCase
}

func NewAuthHandler(authUseCase auth.AuthUseCase) *AuthHandler {
	return &AuthHandler{authUseCase: authUseCase}
}

func (ah *AuthHandler) Auth(c *gin.Context) {
	var a entities.Auth
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{StatusError: ErrInvalidRequestBody})
		return
	}

	if a.Username == "" || a.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username and password are required"})
		return
	}

	token, err := ah.authUseCase.Auth(c.Request.Context(), &a)
	if err != nil {
		if errors.Is(err, ErrInvalidAccessToken) || errors.Is(err, ErrUserDoesNotExist) {
			c.JSON(http.StatusBadRequest, gin.H{StatusError: err.Error()})
			return
		}

		if errors.Is(err, auth.ErrInvalidPassword) {
			c.JSON(http.StatusUnauthorized, gin.H{StatusError: err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{StatusError: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
