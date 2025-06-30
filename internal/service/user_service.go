package service

import (
	"my-gin-app/internal/model"
	"my-gin-app/internal/repository"
)

func GetAllUsers() ([]model.User, error) {
	return repository.FindAllUsers()
}

func AddUser(user *model.User) error {
	return repository.CreateUser(user)
}
