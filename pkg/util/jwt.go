package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(id uint, username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(1 * time.Minute)
	claims := Claims{
		Id: id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer: "yezi",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SigningString()
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, Claims{}, func(t *jwt.Token) (interface{}, error) {
		return nil, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}