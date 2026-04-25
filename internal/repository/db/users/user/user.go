package repository

import (
	"context"
	"encoding/json"

	model "github.com/rchmachina/rach-fw/internal/dto/model/users"
	"github.com/rchmachina/rach-fw/internal/infrastructure/logger"
	"gorm.io/gorm"
)

type IRepositoryUser interface {
	GetWithAddressByID(ctx context.Context, id int64) (*model.UserWithAddress, error)
	GetWithOrdersById(ctx context.Context, id int64) (*model.UserWithOrders, error)
	// CreateUser(ctx context.Context, user *request.CreateUser) error
}

type repositoryUser struct {
	db         *gorm.DB
	SlogLogger logger.Logger
}

func NewUserRepository(db *gorm.DB, logger logger.Logger) IRepositoryUser {
	return &repositoryUser{
		db:         db,
		SlogLogger: logger,
	}
}

func (r *repositoryUser) GetWithAddressByID(ctx context.Context, id int64) (*model.UserWithAddress, error) {
	query := `
	SELECT json_build_object(
		'id', u.id,
		'name', u.name,
		'email', u.email,
		'address', COALESCE(
			json_build_object(
				'street', a.street,
				'city', a.city,
				'province', a.province
			),
			'{}'
		)
	) AS payload
	FROM users."user" u
	LEFT JOIN users.address a ON a.user_id = u.id
	WHERE u.id = ?
	LIMIT 1;
	`
	var payload string
	var user model.UserWithAddress
	err := r.db.WithContext(ctx).
		Raw(query, id).
		Scan(&payload).Error
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(payload), &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repositoryUser) GetWithOrdersById(ctx context.Context, id int64) (*model.UserWithOrders, error) {
	var results model.UserWithOrders

	err := r.db.WithContext(ctx).Raw(`
	SELECT 
		u.id,
		u.name,
		COALESCE(
			json_agg(
				json_build_object(
					'id', o.id,
					'amount', o.amount
				)
			) FILTER (WHERE o.id IS NOT NULL),
			'[]'
		) AS orders
	FROM users.user u
	LEFT JOIN users.orders o ON o.user_id = u.id
	GROUP BY u.id, u.name
	`).Scan(&results).Error

	if err != nil {
		return nil, err
	}
	return &results, nil
}

// func (r *repositoryUser) CreateUser(ctx context.Context, user *request.CreateUser) error {

// 	err := r.db.WithContext(ctx).Exec(`
// 	INSERT INTO users."user" (name, email, password_hash, created_by)
// 	VALUES (?, ?, ?, ?)
// 	`, user.Name, user.Email, user.PasswordHash, user.CreatedBy).Error

// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
