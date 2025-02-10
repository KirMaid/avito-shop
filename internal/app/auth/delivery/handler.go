package delivery

import (
	"avitoshop/internal/app/auth"
	"avitoshop/internal/app/entity"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

const StatusOk = "token"
const StatusError = "errors"

type Response struct {
	Status string `json:"status"`
	Msg    string `json:"message,omitempty"`
}

func newResponse(status string, msg string) *Response {
	return &Response{
		Status: status,
		Msg:    msg,
	}
}

type Handler struct {
	useCase auth.UseCase
}

func newHandler(useCase auth.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

type AuthResponse struct {
	*Response
	Token string `json:"token,omitempty"`
}

func newAuthResponse(status string, msg string, token string) *AuthResponse {
	return &AuthResponse{
		&Response{
			Status: status,
			Msg:    msg,
		},
		token,
	}
}

func (h *Handler) Auth(c *gin.Context) {
	inp := new(entity.User)
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := h.useCase.Auth(c.Request.Context(), inp)
	if err != nil {
		if errors.Is(err, auth.ErrInvalidAccessToken) {
			c.AbortWithStatusJSON(http.StatusBadRequest, newAuthResponse(StatusError, err.Error(), ""))
			return
		}

		if errors.Is(err, auth.ErrUserDoesNotExist) {
			c.AbortWithStatusJSON(http.StatusBadRequest, newAuthResponse(StatusError, err.Error(), ""))
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, newAuthResponse(StatusError, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, newAuthResponse(StatusOk, "", token))
}
