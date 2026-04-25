package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rchmachina/rach-fw/internal/infrastructure/logger"
	"github.com/rchmachina/rach-fw/internal/usecase"
	"github.com/rchmachina/rach-fw/internal/utils/helper"
)

type UserHandler struct {
	uc         usecase.UserUsecase
	SlogLogger logger.Logger
}

func NewUserHandler(uc usecase.UserUsecase, log logger.Logger) *UserHandler {
	return &UserHandler{
		uc:         uc,
		SlogLogger: log,
	}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		h.SlogLogger.With()
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	user, err := h.uc.GetUser(c.Request.Context(), id)
	if err != nil {
		helper.JSONResponse(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	helper.JSONResponse(c, http.StatusOK, user)

}

func (h *UserHandler) GetOrderUser(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	user, err := h.uc.GetOrderByUserid(c.Request.Context(), id)
	if err != nil {
		helper.JSONResponse(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	SuccessResponse(c, "success", user, nil)

}

// func (h *UserHandler) CreateUser(c *gin.Context) {
// 	var user request.CreateUser
// 	var model model.CreateUser
// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	copier.Copy(&user, &model)
// 	err := h.uc.CreateUser(c.Request.Context(), &model)
// 	if err != nil {
// 		ErrorResponse(c, http.StatusInternalServerError, "error when creating user", err)
// 	}

// 	SuccessResponse(c, "success", user, nil)

// }
