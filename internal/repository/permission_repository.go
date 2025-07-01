package repository

import "my-gin-app/internal/model"

func FindPermission(resourceID int, action string) (model.Permission, error) {
	var perm model.Permission
	result := DB.Where("resource_id = ? AND action = ?", resourceID, action).First(&perm)
	return perm, result.Error
}
