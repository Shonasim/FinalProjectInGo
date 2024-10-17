package service

import (
	"FinalProject/internal/models"
	"FinalProject/internal/utils"
	"FinalProject/pkg/errors"
	"fmt"
	"regexp"
)

func (s *Service) GetUserByEmail(email string) error {
	if !isValidEmail(email) {
		return errors.ErrInvalidEmail
	}
	return s.Repository.GetUserByEmail(email)
}

func (s *Service) Registration(user models.User) (*models.User, error) {

	err := validateUser(user)
	if err != nil {
		return nil, err
	}
	hashPass, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, errors.ErrFailedHashing
	}
	user.Password = hashPass
	client, err := s.Repository.AddUser(&user)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (s *Service) SignIn(u *models.User) (string, error) {
	// Проверка существования пользователя
	user, err := s.Repository.GetUser(u.Email)

	// Если ошибка "record not found", то пользователь не существует
	if err != nil {
		return "", err // Вернуть ошибку, если она есть
	}
	if user.Email == "" {
		return "", fmt.Errorf("пользователь с именем %s не найден", u.Email)
	}

	// Проверка пароля пользователя
	if !utils.CheckPasswordHash(u.Password, user.Password) {
		return "", fmt.Errorf("введен неправильный пароль")
	}

	// Генерация JWT токена для авторизованного пользователя
	token, err := utils.GenerateJWT(user)
	if err != nil {
		return "", fmt.Errorf("не удалось сгенерировать токен: %w", err)
	}

	return token, nil
}

func validateUser(user models.User) error {
	if user.Email == "" {
		return errors.ErrInvalidEmail
	}
	if user.Password == "" || len(user.Password) < 8 {
		return errors.ErrInvalidPassword
	}
	return nil
}

func isValidEmail(email string) bool {
	// Регулярное выражение для проверки email
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

/*
Длина от 8 до 20 символов
Содержит хотя бы одну заглавную букву
Содержит хотя бы одну строчную букву
Содержит хотя бы одну цифру
Содержит хотя бы один специальный символ (например, !@#$%^&*)
*/
