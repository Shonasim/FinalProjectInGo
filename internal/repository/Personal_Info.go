package repository

import (
	"FinalProject/internal/models"
	"FinalProject/pkg/errors"
	"time"
)

func (r *Repository) AddPersonalInfo(u *models.PersonalInformation) (*models.PersonalInformation, error) {
	query := `INSERT INTO personal_information 
    (user_id,first_name,last_name,fathers_name,about_me,sex, created_at)
	VALUES (?,?,?,?,?,?,?)`
	err := r.db.Exec(query, u.UserID, u.FirstName, u.LastName, u.FathersName, u.AboutMe, u.Sex, time.Now()).Error
	if err != nil {
		r.logger.Error("Faced an error while tried to insert personal information, err: ", err)
		return nil, errors.ErrFailedCreate
	}
	return u, nil
}
