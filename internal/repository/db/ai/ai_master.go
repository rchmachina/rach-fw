package repository

import (
	"github.com/rchmachina/rach-fw/internal/infrastructure/logger"
	foodRepo "github.com/rchmachina/rach-fw/internal/repository/db/ai/food"
	"gorm.io/gorm"
)

type AiSchema struct {
	FoodRepo   foodRepo.IFoodRepository
	SlogLogger logger.Logger
}

func NewAiSchemaMaster(
	conns *gorm.DB,
	logger logger.Logger,
) AiSchema {
	return AiSchema{
		FoodRepo:   foodRepo.NewFoodRepository(conns),
		SlogLogger: logger,
	}
}
