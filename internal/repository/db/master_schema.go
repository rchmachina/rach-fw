// Package modulerepository is a module for repository
package repositories

import (
	"github.com/rchmachina/rach-fw/internal/infrastructure/logger"

	aiSchema "github.com/rchmachina/rach-fw/internal/repository/db/ai"
	userSchema "github.com/rchmachina/rach-fw/internal/repository/db/users"
	"gorm.io/gorm"
)

type Repositories struct {
	AiSchema   aiSchema.AiSchema
	UserSchema userSchema.UserSchema
	SlogLogger logger.Logger
}

func NewDbRepositoryMaster(
	conns *gorm.DB,
	logger logger.Logger,
) Repositories {
	return Repositories{
		AiSchema:   aiSchema.NewAiSchemaMaster(conns, logger),
		UserSchema: userSchema.NewUserSchemaMaster(conns, logger),
		SlogLogger: logger,
	}
}
