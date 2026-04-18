// Package modulerepository is a module for repository
package repositories

import (
	aiSchema "github.com/rchmachina/rach-fw/internal/repository/ai"
	userSchema "github.com/rchmachina/rach-fw/internal/repository/users"
	"gorm.io/gorm"
)

type Repositories struct {
	AiSchema aiSchema.AiSchema
	UserSchema userSchema.UserSchema
}

func NewRepositoryMaster(
	conns *gorm.DB,
) Repositories {
	return Repositories{
		AiSchema: aiSchema.NewAiSchemaMaster(conns),
		UserSchema:userSchema.NewUserSchemaMaster(conns),
	}
}
