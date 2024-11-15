package util

import (
	"time"
	"todolist/consts"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("JWT_SECRET")

type Claims struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(id uint, username string) (string, error) {
	claims := Claims{
		Id: id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(consts.JWTExpiredTime).Unix(),
			Issuer: "yezi",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (i interface{}, err error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}