package model

import (
	"encoding/json"
	"time"
)

type User struct {
	ID           int64      `gorm:"primaryKey;column:id" json:"id"`
	Name         string     `gorm:"column:name" json:"name"`
	Email        string     `gorm:"column:email" json:"email"`
	PasswordHash string     `gorm:"column:password_hash" json:"password_hash"`
	CreatedAt    time.Time  `gorm:"column:created_at" json:"created_at"`
	CreatedBy    *int64     `gorm:"column:created_by" json:"created_by"`
	UpdatedAt    time.Time  `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy    *int64     `gorm:"column:updated_by" json:"updated_by"`
	DeletedAt    *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	DeletedBy    *int64     `gorm:"column:deleted_by" json:"deleted_by"`
}

type CreateUser struct {
	Name         string    `gorm:"column:name" json:"name"`
	Email        string    `gorm:"column:email" json:"email"`
	PasswordHash string    `gorm:"column:password_hash" json:"password_hash"`
	CreatedBy    *int64    `gorm:"column:created_by" json:"created_by"`
}

type AddressUser struct {
	Street   string `json:"street"`
	City     string `json:"city"`
	Province string `json:"province"`
}

type UserWithAddress struct {
	ID      int64       `json:"id"`
	Name    string      `json:"name"`
	Email   string      `json:"email"`
	Address AddressUser `json:"address"`
}

type UserWithOrders struct {
	ID     int64           `json:"id"`
	Name   string          `json:"name"`
	Orders json.RawMessage `json:"orders"`
}
