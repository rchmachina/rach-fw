package app

import (
	"github.com/rchmachina/rach-fw/configs"
	"github.com/rchmachina/rach-fw/internal/delivery/rest/handler"
	"github.com/rchmachina/rach-fw/internal/delivery/rest/routes"
	"github.com/rchmachina/rach-fw/internal/infrastructure"
	"github.com/rchmachina/rach-fw/internal/infrastructure/logger"
	"github.com/rchmachina/rach-fw/internal/repository"
	"github.com/rchmachina/rach-fw/internal/usecase"
	// "github.com/rchmachina/rach-fw/internal/infrastructure/"
)

type ContainerRest struct {
	UserRoutes routes.RouterUser
	AuthRoutes routes.RouterAuth
	Cfg        configs.Configs
}

func NewContainerRest(configs configs.Configs) *ContainerRest {

	// constanta.RequestIDKey
	db, err := infrastructure.NewDB(configs.DbConf)
	if err != nil {
		panic(err)
	}
	redis, err := infrastructure.NewRedisClient(configs.RedisConf)
	if err != nil {
		panic(err)
	}

	loggerRepo := logger.NewSlogLogger("repository", configs.IsProduction)
	loggerUsecase := logger.NewSlogLogger("usecase", configs.IsProduction)
	loggerHandler := logger.NewSlogLogger("handler", configs.IsProduction)
	loggerRoute := logger.NewSlogLogger("Route", configs.IsProduction)

	repository := repository.NewMasterRepo(db, redis, loggerRepo)

	userUsecase := usecase.NewUserUsecase(repository, loggerUsecase)
	userHandler := handler.NewUserHandler(userUsecase, loggerHandler)
	userRoutes := routes.NewRouterUser(userHandler, loggerRoute, configs)

	authUsecase := usecase.NewAuthUserCase(repository, loggerUsecase, configs)
	AuthHandler := handler.NewAuthHandler(authUsecase, loggerHandler)
	AuthRoutes := routes.NewAuthRouter(AuthHandler)

	return &ContainerRest{
		UserRoutes: *userRoutes,
		AuthRoutes: *AuthRoutes,
		Cfg:        configs,
	}
}
