package http

import (
	usecase "avitoshop/internal/app/usecases/user_info"
	"github.com/gin-gonic/gin"
)

type UserInfoHandler struct {
	userInfoUseCase usecase.UserInfoUseCase
}

func NewUserInfoHandler(userInfoUseCase usecase.UserInfoUseCase) *UserInfoHandler {
	return &UserInfoHandler{userInfoUseCase: userInfoUseCase}
}

func (uih *UserInfoHandler) GetInfo(c *gin.Context) {

}
