package model

import jwt "github.com/golang-jwt/jwt/v5"

type TokenValue struct {
	Email string `json:"email"`
	Id    string `json:"id"`
	Name  string `json:"name"`

	jwt.RegisteredClaims
}