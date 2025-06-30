package service

import (
	"my-gin-app/internal/model"
	"my-gin-app/internal/repository"
)

func GetAllDepartments() ([]model.Department, error) {
	return repository.FindAllDepartments()
}

func AddDepartment(dept *model.Department) error {
	return repository.CreateDepartment(dept)
}
