package repository

import "my-gin-app/internal/model"

func FindAllDepartments() ([]model.Department, error) {
	var departments []model.Department
	result := DB.Find(&departments)
	return departments, result.Error
}

func CreateDepartment(dept *model.Department) error {
	return DB.Create(dept).Error
}
