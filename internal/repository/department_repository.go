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

func FindDepartmentByID(id int) (model.Department, error) {
	var dept model.Department
	result := DB.Where("id = ?", id).First(&dept)
	return dept, result.Error
}

func UpdateDepartment(dept *model.Department) error {
	return DB.Save(dept).Error
}

func DeleteDepartment(id int) error {
	return DB.Delete(&model.Department{}, id).Error
}
