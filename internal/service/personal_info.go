package service

import (
	"FinalProject/internal/models"
	"FinalProject/pkg/errors"
)

func (s *Service) FillExtraInfo(information models.PersonalInformation) (*models.PersonalInformation, error) {
	err := validatePersonaInfo(information)
	if err != nil {
		return nil, err
	}
	return s.Repository.AddPersonalInfo(&information)
}

func (s *Service) GetPersInfoById(userId int) (*models.PersonalInformation, error) {
	return s.Repository.GetPersonalInfoById(userId)
}

func (s *Service) UploadPhoto(userId int, path string) error {
	return s.Repository.UploadPhoto(userId, path)
}

func validatePersonaInfo(user models.PersonalInformation) error {
	if user.FirstName == "" {
		return errors.ErrInvalidFirstName
	}
	if user.LastName == "" {
		return errors.ErrInvalidLastName
	}
	if user.FathersName == "" {
		return errors.ErrInvalidFathersName
	}
	if user.AboutMe == "" {
		return errors.ErrInvalidAboutUser
	}
	if user.Sex == "" {
		return errors.ErrInvalidSex
	}
	return nil
}
