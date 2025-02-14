package http

import (
	userinfo "avitoshop/internal/app/usecases/user_info"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfoHandler struct {
	userInfoUseCase userinfo.UserInfoUseCase
}

func NewUserInfoHandler(userInfoUseCase userinfo.UserInfoUseCase) *UserInfoHandler {
	return &UserInfoHandler{userInfoUseCase: userInfoUseCase}
}

func (uih *UserInfoHandler) GetInfo(c *gin.Context) {
	usernameInterface, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": "username not found in context"})
		return
	}

	username, ok := usernameInterface.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": "username is not a string"})
		return
	}

	userInfoDTO, err := uih.userInfoUseCase.GetInfo(c.Request.Context(), username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userInfoDTO)
}
