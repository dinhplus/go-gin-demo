package repository

import "my-gin-app/internal/model"

func FindRoleByID(id int) (model.Role, error) {
	var role model.Role
	result := DB.Preload("Permissions").First(&role, id)
	return role, result.Error
}

func FindAllRoles() ([]model.Role, error) {
	var roles []model.Role
	result := DB.Preload("Permissions").Find(&roles)
	return roles, result.Error
}

func CreateRole(role *model.Role) error {
	return DB.Create(role).Error
}

func UpdateRole(role *model.Role) error {
	return DB.Save(role).Error
}

func DeleteRole(id int) error {
	return DB.Delete(&model.Role{}, id).Error
}
