package service

import "github.com/dgrijalva/jwt-go"

type JWTService interface {
	GenerateToken(userId uint64, Username string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaims struct {
	UserId uint64 `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}