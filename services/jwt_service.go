package services

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"time"
)

var jwtSecret = []byte("secret_key")

type Claims struct {
	AccountID string   `json:"account_id"`
	jti    string `json:"jti"`
	jwt.StandardClaims
}

func GenerateToken(accountID string) (string, error) {
	expirationTime := time.Now().Add(24 * 30 * time.Hour) // 设置30天有效期

	claims := &Claims{
		AccountID: accountID,
		jti:    uuid.New().String(), // 生成唯一的 JWT ID
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Id:        uuid.New().String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return claims, nil
}
