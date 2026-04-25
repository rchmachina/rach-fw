package repository

import (
	"github.com/rchmachina/rach-fw/internal/infrastructure/logger"
	usersRepo "github.com/rchmachina/rach-fw/internal/repository/db/users/user"
	"gorm.io/gorm"
)

type UserSchema struct {
	UsersRepo usersRepo.IRepositoryUser
	AuthRepo  usersRepo.IRepositoryAuth
}

func NewUserSchemaMaster(
	conns *gorm.DB,
	log logger.Logger,
) UserSchema {
	return UserSchema{
		UsersRepo: usersRepo.NewUserRepository(conns, log),
		AuthRepo:  usersRepo.NewAuthRepository(conns, log),
	}
}
