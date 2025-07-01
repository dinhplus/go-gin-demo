package repository

import "my-gin-app/internal/model"

func FindPermission(resourceID int, action string) (model.Permission, error) {
	var perm model.Permission
	result := DB.Where("resource_id = ? AND action = ?", resourceID, action).First(&perm)
	return perm, result.Error
}

func FindAllPermissions() ([]model.Permission, error) {
	var perms []model.Permission
	result := DB.Find(&perms)
	return perms, result.Error
}

func CreatePermission(perm *model.Permission) error {
	return DB.Create(perm).Error
}
