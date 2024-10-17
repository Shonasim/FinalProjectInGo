package repository

import (
	"FinalProject/internal/models"
	errors2 "FinalProject/pkg/errors"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"time"
)

func (r *Repository) AddUser(u *models.User) (*models.User, error) {
	query := `INSERT INTO users
    (email,password,created_at) 
    VALUES (?,?,?)`
	err := r.db.Exec(query, u.Email, u.Password, time.Now())
	if err != nil {
		r.logger.Error("Faced an error while tried to insert user, err: ", err)
		return nil, errors2.ErrFailedCreate
	}
	return u, nil
}

func (r *Repository) GetUserByID(id int) (*models.User, error) {
	var user models.User

	// select * from users where user_id = id
	result := r.db.First(&user, id)
	if result.Error != nil {
		log.Printf("GetUserByID: Failed to get user: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to get user: %v\n", result.Error)
	}
	return &user, nil
}
func (r *Repository) GetUser(email string) (models.User, error) {
	var user models.User
	sql := `select * from users where email = ?;`
	err := r.db.Raw(sql, email).Scan(&user).Error
	if err != nil {
		r.logger.Error("Faced an error while tried to select user from table, err: ", err)
		return models.User{}, err
	}
	return user, nil
}
func (r *Repository) GetUserByEmail(email string) error {
	var user models.User
	sql := `select * from users where email = ?;`
	err := r.db.Raw(sql, email).Scan(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		} else {
			r.logger.Error("Faced an error while tried to select user from table, err: ", err)
			return err
		}
	}
	if user.Email == "" {
		return nil
	}
	return errors2.ErrAlreadyExists
}

func (r *Repository) UpdateUser(u *models.User) (*models.User, error) {
	// update users set username = 'admin', password = 'admin' where user_id = 1
	result := r.db.Model(&u).Clauses(clause.Returning{}).Updates(&u)
	if result.Error != nil {
		log.Printf("UpdateUser: Failed to update user: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to update user: %v\n", result.Error)
	}

	return u, nil
}

func (r *Repository) DeleteUser(id int) (int, error) {
	// delete from users where user_id = id
	result := r.db.Delete(&models.User{}, id)
	if result.Error != nil {
		log.Printf("DeleteUser: Failed to delete user: %v\n", result.Error)
		return 0, fmt.Errorf("Failed to delete user: %v\n", result.Error)
	}

	return id, nil
}
