package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("RAHASIA_NEGARA_123")

type Claims struct {
	UserID     uint   `json:"user_id"`
	Role       string `json:"role"`
	Department string `json:"department"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint, role, department string) (string, error) {
	claims := Claims{
		UserID:     userID,
		Role:       role,
		Department: department,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("token tidak valid")
	}

	return claims, nil
}
