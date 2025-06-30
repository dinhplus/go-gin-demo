package repository

import "my-gin-app/internal/model"

func FindPermission(resource, action string) (model.Permission, error) {
	var perm model.Permission
	result := DB.Where("resource = ? AND action = ?", resource, action).First(&perm)
	return perm, result.Error
}
