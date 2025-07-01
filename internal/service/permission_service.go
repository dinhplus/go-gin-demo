package service

import (
	"my-gin-app/internal/model"
	"my-gin-app/internal/repository"
)

func GetPermission(resourceID int, action string) (model.Permission, error) {
	return repository.FindPermission(resourceID, action)
}

func GetAllPermissions() ([]model.Permission, error) {
	return repository.FindAllPermissions()
}

func AddPermission(perm *model.Permission) error {
	return repository.CreatePermission(perm)
}
