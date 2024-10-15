package service

import (
	"FinalProject/internal/models"
	"FinalProject/internal/utils"
	"FinalProject/pkg/errors"
	errors2 "errors"
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
	return s.Repository.AddUser(&user)
}

// ListUsers функция, которая возвращает список всех пользователей из репозитория.
func (s *Service) ListUsers() ([]models.User, error) {
	// Получаем список пользователей из репозитория
	users, err := s.Repository.GetUsers()
	if err != nil {
		return nil, err
	}

	// Проверяем, что список пользователей не пустой
	if len(users) == 0 {
		return nil, fmt.Errorf("no users found")
	}

	return users, nil
}

// FindUser функция, которая возвращает пользователя по ID.
func (s *Service) FindUser(id int) (*models.User, error) {
	// Получаем пользователя по ID из репозитория
	userByID, err := s.Repository.GetUserByID(id)
	if err != nil {
		// Если произошла ошибка и это ошибка "record not found", возвращаем её
		if errors2.As(err, errors.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with id %d not found", id)
		}
		return nil, err
	}

	return userByID, nil
}

// EditUser функция, которая редактирует пользователя в репозитории.
func (s *Service) EditUser(u *models.User) (*models.User, error) {
	// Получаем пользователя по ID из репозитория
	_, err := s.Repository.GetUserByID(u.UserID)
	if err != nil {
		// Если произошла ошибка и это ошибка "record not found", возвращаем её
		if errors2.As(err, errors.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with id %d not found", u.UserID)
		}
		return nil, err
	}

	// Обновляем информацию о пользователе
	updatedUser, err := s.Repository.UpdateUser(u)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	// Не отправляем пароль в ответ
	if updatedUser.Password != "" {
		updatedUser.Password = ""
	}

	return updatedUser, nil
}

// DeleteUser функция, которая удаляет пользователя из репозитория.
func (s *Service) DeleteUser(id int) (int, error) {
	// Получаем пользователя по ID из репозитория
	_, err := s.Repository.GetUserByID(id)
	if err != nil {
		// Если произошла ошибка и это ошибка "record not found", возвращаем её
		if errors2.Is(err, errors.ErrRecordNotFound) {
			return 0, fmt.Errorf("user with id %d not found", id)
		}
		return 0, err
	}

	// Удаляем пользователя
	deletedRows, err := s.Repository.DeleteUser(id)
	if err != nil {
		return 0, fmt.Errorf("failed to delete user with id %d: %w", id, err)
	}

	return deletedRows, nil
}

func validateUser(user models.User) error {
	//if user.FirstName == "" {
	//	return errors.ErrInvalidFirstName
	//}
	//if user.LastName == "" {
	//	return errors.ErrInvalidLastName
	//}
	//if user.FathersName == "" {
	//	return errors.ErrInvalidFathersName
	//}
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
