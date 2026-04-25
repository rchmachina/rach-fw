package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rchmachina/rach-fw/configs"
	"github.com/rchmachina/rach-fw/internal/delivery/rest/handler"
	"github.com/rchmachina/rach-fw/internal/delivery/rest/middleware"
	"github.com/rchmachina/rach-fw/internal/infrastructure/logger"
)

type RouterUser struct {
	handler *handler.UserHandler
	logger  logger.Logger
	cfg     configs.Configs
}

func NewRouterUser(h *handler.UserHandler, l logger.Logger, c configs.Configs) *RouterUser {
	return &RouterUser{
		handler: h,
		logger:  l,
		cfg:     c,
	}
}

func (r *RouterUser) Register(parent *gin.RouterGroup) {
	usersV1 := parent.Group("/v1/users")

	usersV1.Use(
		middleware.Auth(r.logger, r.cfg.AccesTokenKey),
	)
	usersV1.GET("/:id", r.handler.GetUser)
	usersV1.GET("/:id/order", r.handler.GetOrderUser)
}
