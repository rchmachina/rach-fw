package jwt

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	constant "github.com/rchmachina/rach-fw/internal/const"
	model "github.com/rchmachina/rach-fw/internal/dto/model/jwt"
)

func GenerateAccessToken(data model.TokenValue, secretKey string, expTime time.Duration) (string, error) {

	claims := model.TokenValue{
		Email: data.Email,
		Id:    data.Id,
		Name:  data.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretKey))
}

func GenerateRefreshToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func VerifyToken(tokenString string, secretKey string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func DecodeToken(tokenStr, secret string) (*model.TokenValue, error) {

	token, err := jwt.ParseWithClaims(tokenStr, &model.TokenValue{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*model.TokenValue)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

// // Setter
// func SetUserInfo(ctx context.Context, val map[string]interface{}) context.Context {
// 	return context.WithValue(ctx, constant.UserInfoKey, val)
// }

// Getter (safe)
func GetUserInfo(ctx context.Context) (*model.TokenValue, error) {
	val := ctx.Value(constant.UserInfoKey)
	if val == nil {
		return nil, fmt.Errorf("user info not found")
	}

	user, ok := val.(*model.TokenValue)
	if !ok {
		return nil, fmt.Errorf("invalid user info type")
	}

	return user, nil
}
