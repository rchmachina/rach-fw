package http

import (
	"github.com/gin-gonic/gin"
	"github.com/rchmachina/rach-fw/internal/app"
	"github.com/rchmachina/rach-fw/internal/delivery/rest/middleware"
	userRouter "github.com/rchmachina/rach-fw/internal/delivery/rest/routes"
)

type Router struct {
    userRouter *userRouter.RouterUser
}

func NewRouter(router *app.ContainerUser) *Router {
    return &Router{
        userRouter: &router.UserRoutes,
    }
}

func (r *Router) Setup() *gin.Engine {
    router := gin.Default()

    router.Use(
        middleware.Logger(),
        middleware.RequestIDMiddleware(),
    )

    api := router.Group("/api")
    r.userRouter.Register(api)

    return router
}