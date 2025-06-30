package repository

import "my-gin-app/internal/model"

func FindRoleByID(id int) (model.Role, error) {
	var role model.Role
	result := DB.Preload("Permissions").First(&role, id)
	return role, result.Error
}
