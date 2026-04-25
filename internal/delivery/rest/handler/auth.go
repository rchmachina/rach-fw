package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	request "github.com/rchmachina/rach-fw/internal/dto/request/auth"
	"github.com/rchmachina/rach-fw/internal/infrastructure/logger"
	"github.com/rchmachina/rach-fw/internal/usecase"
)

type AuthHandler struct {
	uc         usecase.AuthUsecase
	SlogLogger logger.Logger
}

func NewAuthHandler(uc usecase.AuthUsecase, log logger.Logger) *AuthHandler {
	return &AuthHandler{
		uc:         uc,
		SlogLogger: log,
	}
}

func (h *AuthHandler) LoginUser(c *gin.Context) {
	var user request.LoginUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := h.uc.LoginAuth(c.Request.Context(), &user)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "error when creating user", err)
		return
	}

	SuccessResponse(c, "success", token, nil)

}

func (h *AuthHandler) CreateUser(c *gin.Context) {
	var user request.CreateAuth
	if err := c.ShouldBindJSON(&user); err != nil {
		ErrorResponse(c, http.StatusBadRequest, "error when creating user", err)
		return
	}
	token, err := h.uc.CreateUser(c.Request.Context(), user)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "error when creating user", err)
		return
	}

	SuccessResponse(c, "success", token, nil)

}

func (h *AuthHandler) LogOutPerdevice(c *gin.Context) {
	var token request.Token
	if err := c.ShouldBindJSON(&token); err != nil {
		ErrorResponse(c, http.StatusBadRequest, "error when creating user", err)
		return
	}

	err := h.uc.RevokeAuth(c.Request.Context(), token)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "error when log out", err)
		return
	}

	SuccessResponse(c, "success log out ", nil, nil)

}

func (h *AuthHandler) GenNewAccessToken(c *gin.Context) {
	var token request.Token
	if err := c.ShouldBindJSON(&token); err != nil {
		ErrorResponse(c, http.StatusBadRequest, "error when creating user", err)
		return
	}

	key, err := h.uc.CreateNewAccessToken(c.Request.Context(), token)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "error when log out", err)
		return
	}

	SuccessResponse(c, "success refresh token",key, nil)

}
