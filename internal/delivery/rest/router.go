package http

import (
	"github.com/gin-gonic/gin"
	"github.com/rchmachina/rach-fw/configs"
	"github.com/rchmachina/rach-fw/internal/app"
	"github.com/rchmachina/rach-fw/internal/delivery/rest/middleware"
	routers "github.com/rchmachina/rach-fw/internal/delivery/rest/routes"

	"github.com/rchmachina/rach-fw/internal/infrastructure/logger"
)

type Router struct {
	userRouter *routers.RouterUser
	authRouter *routers.RouterAuth
	cfg        configs.Configs
}

func NewRouter(container *app.ContainerRest) *Router {
	return &Router{
		userRouter: &container.UserRoutes,
		authRouter: &container.AuthRoutes,
		cfg:        container.Cfg,
	}
}

func (r *Router) Setup() *gin.Engine {
	router := gin.Default()

	loggerMiddlewareRest := logger.NewSlogLogger("middlewareRest", r.cfg.IsProduction)

	router.Use(
		middleware.IncomingRequest(loggerMiddlewareRest),
		middleware.RequestIDMiddleware(),
	)

	api := router.Group("/api")
	r.userRouter.Register(api)
	r.authRouter.Register(api,loggerMiddlewareRest,r.cfg)

	return router
}
