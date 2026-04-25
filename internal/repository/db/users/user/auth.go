package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	model "github.com/rchmachina/rach-fw/internal/dto/model/auth"
	request "github.com/rchmachina/rach-fw/internal/dto/request/auth"
	"github.com/rchmachina/rach-fw/internal/infrastructure/logger"
	"gorm.io/gorm"
)

type IRepositoryAuth interface {
	GetUserByEmail(ctx context.Context, email string) (model.GetLoginUser, error)
	CreateUser(ctx context.Context, req request.CreateAuth) (string, error)
	// CreateUser(ctx context.Context, user *request.CreateUser) error
}

type repositoryAuth struct {
	db         *gorm.DB
	SlogLogger logger.Logger
}

func NewAuthRepository(db *gorm.DB, logger logger.Logger) IRepositoryAuth {
	return &repositoryAuth{
		db:         db,
		SlogLogger: logger,
	}
}

func (r *repositoryAuth) GetUserByEmail(ctx context.Context, email string) (model.GetLoginUser, error) {
	var results model.GetLoginUser

	err := r.db.WithContext(ctx).Raw(`
	SELECT 
		u.id,
		u.email,
		u.name,
		u.password_hash
	FROM users.user u
	where u.email = ?


	`, email).Scan(&results).Error
	if err != nil {
		return results, err
	}
	if results.Email == "" {
		return results, fmt.Errorf("email not exist")
	}
	return results, nil
}

func (r *repositoryAuth) CreateUser(ctx context.Context, req request.CreateAuth) (string, error) {
	query := `
		INSERT INTO users."user" (name, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	var id string

	err := r.db.WithContext(ctx).
		Raw(query, req.Name, req.Email, req.PasswordHash).
		Scan(&id).Error

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return "", errors.New("email already registered")
			}
		}
		return "", err
	}

	return id, nil
}
