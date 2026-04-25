package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rchmachina/rach-fw/configs"
	"github.com/rchmachina/rach-fw/internal/delivery/rest/handler"
	"github.com/rchmachina/rach-fw/internal/delivery/rest/middleware"
	"github.com/rchmachina/rach-fw/internal/infrastructure/logger"
)

type RouterAuth struct {
	handler *handler.AuthHandler
}

func NewAuthRouter(h *handler.AuthHandler) *RouterAuth {
	return &RouterAuth{
		handler: h,
	}
}

func (r *RouterAuth) Register(parent *gin.RouterGroup, log logger.Logger, cfg configs.Configs) {
	authV1 := parent.Group("/v1/auth")

	// public
	authV1.POST("/login", r.handler.LoginUser)
	authV1.POST("/sign-up", r.handler.CreateUser)
	authV1.POST("/gen-access-token", r.handler.GenNewAccessToken)

	// private
	authProtected := authV1.Group("")
	authProtected.Use(middleware.Auth(log, cfg.AccesTokenKey))

	authProtected.POST("/log-out", r.handler.LogOutPerdevice)

}
