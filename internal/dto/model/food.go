package model

import "github.com/rchmachina/rach-fw/internal/utils/helper"

type Food struct {
	ID          int64     `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Category    string    `json:"category"`
	Calories    float64   `json:"calories"`
	Protein     float64   `json:"protein"`
	Fat         float64   `json:"fat"`
	Carbs       float64   `json:"carbs"`
	Description string    `json:"description"`

	Embedding   []float32 `gorm:"-" json:"embedding"`
}

func (f *Food) GetEmbeddingString()string{
	return helper.VectorToString(f.Embedding)
}