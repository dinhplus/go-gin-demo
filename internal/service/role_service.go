package service

import (
	"my-gin-app/internal/model"
	"my-gin-app/internal/repository"
)

func GetRoleByID(id int) (model.Role, error) {
	return repository.FindRoleByID(id)
}
