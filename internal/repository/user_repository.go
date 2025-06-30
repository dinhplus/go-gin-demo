package repository

import "my-gin-app/internal/model"

func FindAllUsers() ([]model.User, error) {
	var users []model.User
	result := DB.Preload("Department").Preload("Position").Preload("Stack").Find(&users)
	return users, result.Error
}

func CreateUser(user *model.User) error {
	return DB.Create(user).Error
}
