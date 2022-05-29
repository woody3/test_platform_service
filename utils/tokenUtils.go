package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const TokenExpireDuration = time.Minute * 30

var authSecret []byte

type CustomClaim struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenToken(username string) (string, error) {
	authSecret = []byte(GetConfig().GetString("authSecret"))
	claims := CustomClaim{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    "test-platform",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(authSecret)
}

func ParseToken(tokenString string) (*CustomClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaim{}, func(token *jwt.Token) (i interface{}, err error) {
		return authSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// 校验token
	if claims, ok := token.Claims.(*CustomClaim); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
