package service

import (
	"grace/database"
	"grace/model"
)

type UserService struct{}

func (s *UserService) CreateUser(user *model.User) error {
	return database.GetDB().Create(user).Error
}

func (s *UserService) GetUsers() ([]model.User, error) {
	var users []model.User
	err := database.GetDB().Find(&users).Error
	return users, err
}

func (s *UserService) GetUserByID(id uint) (*model.User, error) {
	var user model.User
	err := database.GetDB().First(&user, id).Error
	return &user, err
}

func (s *UserService) UpdateUser(user *model.User) error {
	return database.GetDB().Save(user).Error
}

func (s *UserService) DeleteUser(id uint) error {
	return database.GetDB().Delete(&model.User{}, id).Error
}
