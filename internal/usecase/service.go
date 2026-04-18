package usecase

import (
	"context"
	"errors"

	model "github.com/rchmachina/rach-fw/internal/dto/model/users"
	repo "github.com/rchmachina/rach-fw/internal/repository"
)

type UserUsecase interface {
	GetUser(ctx context.Context, id int64) (*model.UserWithAddress, error)
	GetOrderByUserid(ctx context.Context, id int64) (*model.UserWithOrders, error)
	CreateUser(ctx context.Context, user *model.CreateUser) error
}

type userUsecase struct {
	repo repo.Repositories
}

func NewUserUsecase(r repo.Repositories) UserUsecase {
	return &userUsecase{repo: r}
}

func (u *userUsecase) GetUser(ctx context.Context, id int64) (*model.UserWithAddress, error) {
	if id <= 0 {
		return nil, errors.New("invalid id")
	}

	data, err := u.repo.UserSchema.UsersRepo.GetWithAddressByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return data, nil
}
func (u *userUsecase) GetOrderByUserid(ctx context.Context, id int64) (*model.UserWithOrders, error) {
	if id <= 0 {
		return nil, errors.New("invalid id")
	}

	data, err := u.repo.UserSchema.UsersRepo.GetWithOrdersById(ctx, id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *userUsecase) CreateUser(ctx context.Context, user *model.CreateUser) error {

	return u.repo.UserSchema.UsersRepo.CreateUser(ctx, user)
}
