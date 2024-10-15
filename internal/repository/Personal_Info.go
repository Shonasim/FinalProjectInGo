package repository

import (
	"FinalProject/internal/models"
	errors2 "FinalProject/pkg/errors"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
)

func (r *Repository) AddPersonalInfo(u *models.PersonalInformation) (*models.PersonalInformation, error) {
	result := r.db.Create(&u)
	if result.Error != nil {
		return nil, fmt.Errorf("Failed to add user: %v\n", result.Error)
	}
	return u, nil
}

func (r *Repository) GetPersonalInfoUsers() ([]models.PersonalInformation, error) {
	var users []models.PersonalInformation

	// select * from users
	result := r.db.Find(&users)
	if result.Error != nil {
		log.Printf("GetUsers: Failed to get users: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to get users: %v\n", result.Error)
	}
	return users, nil
}
func (r *Repository) GetPersonalInfoById(id int) (*models.PersonalInformation, error) {
	var user models.PersonalInformation

	// select * from users where user_id = id
	result := r.db.First(&user, id)
	if result.Error != nil {
		log.Printf("GetUserByID: Failed to get user: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to get user: %v\n", result.Error)
	}
	return &user, nil
}

func (r *Repository) GetUserByLastName(lastName string) (*models.PersonalInformation, error) {
	var user models.PersonalInformation
	sql := `select * from personal_information where last_name = ?;`
	err := r.db.Raw(sql, lastName).Scan(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			r.logger.Error("Faced an error while tried to select user from table, err: ", err)
			return nil, nil
		}
	}
	if user.LastName == "" {
		return nil, err
	}
	return &user, errors2.ErrAlreadyExists
}

func (r *Repository) UpdatePersonalInfo(u *models.PersonalInformation) (*models.PersonalInformation, error) {
	// update users set username = 'admin', password = 'admin' where user_id = 1
	result := r.db.Model(&u).Clauses(clause.Returning{}).Updates(&u)
	if result.Error != nil {
		log.Printf("UpdateUser: Failed to update user: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to update user: %v\n", result.Error)
	}

	return u, nil
}
