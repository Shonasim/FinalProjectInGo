package handlers

//
//import (
//	"FinalProject/internal/db"
//	"FinalProject/internal/models"
//	"encoding/json"
//	"github.com/dgrijalva/jwt-go"
//	"net/http"
//	"time"
//)
//
//func RegisterHandler(w http.ResponseWriter, r *http.Request) {
//	var user models.User
//	err := json.NewDecoder(r.Body).Decode(&user)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	// Проверка, существует ли пользователь с таким email
//	var existingUser models.User
//	if err := db.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
//		http.Error(w, "User already exists", http.StatusConflict)
//		return
//	}
//
//	user.Active = true
//
//	// Сохранение нового пользователя в БД
//	if err := db.DB.Create(&user).Error; err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	w.WriteHeader(http.StatusCreated)
//}
//
//func LoginHandler(w http.ResponseWriter, r *http.Request) {
//	var user models.User
//	err := json.NewDecoder(r.Body).Decode(&user)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	var dbUser models.User
//	if err := db.DB.Where("email = ?", user.Email).First(&dbUser).Error; err != nil {
//		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
//		return
//	}
//
//	if user.Password != dbUser.Password {
//		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
//		return
//	}
//
//	claims := models.JWTClaim{
//		UserID: dbUser.ID,
//		StandardClaims: jwt.StandardClaims{
//			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
//		},
//	}
//
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	tokenString, err := token.SignedString([]byte("your_secret_key"))
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	w.Header().Set("Authorization", "Bearer "+tokenString)
//	w.WriteHeader(http.StatusOK)
//}
