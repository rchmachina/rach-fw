package repository

import (
	usersRepo "github.com/rchmachina/rach-fw/internal/repository/users/user"
	"gorm.io/gorm"
)

type UserSchema struct {
	UsersRepo usersRepo.IRepositoryUser
}

func NewUserSchemaMaster(
	conns *gorm.DB,
) UserSchema {
	return UserSchema{
		UsersRepo: usersRepo.NewUserRepository(conns),
	}
}
