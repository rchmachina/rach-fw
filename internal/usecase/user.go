package usecase

import (
	"context"
	"errors"

	model "github.com/rchmachina/rach-fw/internal/dto/model/users"
	logger "github.com/rchmachina/rach-fw/internal/infrastructure/logger"
	repo "github.com/rchmachina/rach-fw/internal/repository"
)

type UserUsecase interface {
	GetUser(ctx context.Context, id int64) (*model.UserWithAddress, error)
	GetOrderByUserid(ctx context.Context, id int64) (*model.UserWithOrders, error)
}

type userUsecase struct {
	repo   repo.Repository
	logger logger.Logger
}

func NewUserUsecase(r *repo.Repository, logger logger.Logger) UserUsecase {
	return &userUsecase{
		repo:   *r,
		logger: logger,
	}
}

func (u *userUsecase) GetUser(ctx context.Context, id int64) (*model.UserWithAddress, error) {
	if id <= 0 {
		u.logger.Error("invalid id provided", logger.Field{Key: "id", Value: id})
		return nil, errors.New("invalid id")
	}

	data, err := u.repo.Sql.UserSchema.UsersRepo.GetWithAddressByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return data, nil
}
func (u *userUsecase) GetOrderByUserid(ctx context.Context, id int64) (*model.UserWithOrders, error) {
	if id <= 0 {
		return nil, errors.New("invalid id")
	}

	data, err := u.repo.Sql.UserSchema.UsersRepo.GetWithOrdersById(ctx, id)
	if err != nil {
		return nil, err
	}

	return data, nil
}
