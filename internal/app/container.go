package app

import (
	"github.com/rchmachina/rach-fw/internal/delivery/rest/handler"
	"github.com/rchmachina/rach-fw/internal/delivery/rest/routes"
	repo "github.com/rchmachina/rach-fw/internal/repository"
	"github.com/rchmachina/rach-fw/internal/usecase"
	"gorm.io/gorm"
)

type ContainerUser struct {
    UserRoutes routes.RouterUser
}

func NewContainerUser(db  *gorm.DB) *ContainerUser {
    userRepo := repo.NewRepositoryMaster(db)
    userUsecase := usecase.NewUserUsecase(userRepo)
    userHandler := handler.NewUserHandler(userUsecase)
	userRoutes := routes.NewRouterUser(userHandler)

    return &ContainerUser{
        UserRoutes: *userRoutes,
    }
}