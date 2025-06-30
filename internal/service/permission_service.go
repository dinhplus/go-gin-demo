package service

import (
	"my-gin-app/internal/model"
	"my-gin-app/internal/repository"
)

func GetPermission(resource, action string) (model.Permission, error) {
	return repository.FindPermission(resource, action)
}
