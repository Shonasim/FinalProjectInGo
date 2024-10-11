package models

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type User struct {
	UserID      int       `json:"user_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	FathersName string    `json:"fathers_name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	RoleID      int       `json:"role_id"`
	CreatedAt   time.Time `json:"created_at"`
	Active      bool      `json:"active"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type JWTClaim struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}
