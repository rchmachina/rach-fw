package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/rchmachina/rach-fw/internal/delivery/rest/handler"
)

type RouterUser struct {
    handler *handler.UserHandler
}

func NewRouterUser(h *handler.UserHandler) *RouterUser {
    return &RouterUser{handler: h}
}

func (r *RouterUser) Register(parent *gin.RouterGroup) {
    usersV1 := parent.Group("/v1/users")
    usersV1.GET("/:id", r.handler.GetUser)
    usersV1.GET("/:id/order", r.handler.GetOrderUser)
}