package models

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type User struct {
	UserID    int       `json:"user_id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	Active    bool      `json:"active"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type UserCreateResponse struct {
	Email string `json:"email"`
}

type JWTClaim struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}
