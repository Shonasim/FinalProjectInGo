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

type PersonalInformation struct {
	InfoID      int       `json:"info_id"`
	UserID      int       `json:"user_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	FathersName string    `json:"fathers_name"`
	AboutMe     string    `json:"about_me"`
	Sex         string    `json:"sex"`
	Photo       string    `json:"photo"`
	CreatedAt   time.Time `json:"created_at"`
	Active      bool      `json:"active"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type UserCreateResponse struct {
	Email string `json:"email"`
}

type JWTClaim struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}
