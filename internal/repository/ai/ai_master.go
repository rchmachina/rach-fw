package repository

import (
	foodRepo "github.com/rchmachina/rach-fw/internal/repository/ai/food"
	"gorm.io/gorm"
)

type AiSchema struct {
	FoodRepo foodRepo.IFoodRepository
}

func NewAiSchemaMaster(
	conns *gorm.DB,
) AiSchema {
	return AiSchema{
		FoodRepo: foodRepo.NewFoodRepository(conns),
	}
}
