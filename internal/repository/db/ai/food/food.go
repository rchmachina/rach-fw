package repository

import (
	"context"
	"github.com/rchmachina/rach-fw/internal/dto/model"
	"github.com/rchmachina/rach-fw/internal/utils/helper"
	"gorm.io/gorm"
)



type foodRepository struct {
	db *gorm.DB
}

func NewFoodRepository(db *gorm.DB) IFoodRepository {
	return &foodRepository{db: db}
	
}

type IFoodRepository interface {
	Insert(ctx context.Context, food *model.Food) error
	SearchSimilar(ctx context.Context, embedding []float32, limit int) ([]model.Food, error)
}


func (r *foodRepository) Insert(ctx context.Context, food *model.Food) error {
	

	
	query := `
		INSERT INTO ai.foods 
		(name, category, calories, protein, fat, carbs, description, embedding)
		VALUES (?,?,?,?,?,?,?,?)
	`

	return r.db.WithContext(ctx).Exec(query,
		food.Name,
		food.Category,
		food.Calories,
		food.Protein,
		food.Fat,
		food.Carbs,
		food.Description,
		food.GetEmbeddingString(),
	).Error
}


func (r *foodRepository) SearchSimilar(ctx context.Context, embedding []float32, limit int) ([]model.Food, error) {

	query := `
		SELECT id, name, category, calories, protein, fat, carbs, description
		FROM ai.foods
		ORDER BY embedding <=> ?
		LIMIT ?
	`

	var foods []model.Food

	err := r.db.WithContext(ctx).
		Raw(query, helper.VectorToString(embedding), limit).
		Scan(&foods).Error

	return foods, err
}