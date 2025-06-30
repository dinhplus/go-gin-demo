package repository

import "my-gin-app/internal/model"

func SeedRoles() {
	var count int64
	DB.Model(&model.Role{}).Count(&count)
	if count == 0 {
		roles := []model.Role{
			{ID: 1, Name: "admin"},
			{ID: 2, Name: "user"},
		}
		for _, r := range roles {
			DB.Create(&r)
		}
	}
}

func SeedDepartments() {
	var count int64
	DB.Model(&model.Department{}).Count(&count)
	if count == 0 {
		departments := []model.Department{
			{ID: 1, Name: "Default Department"},
		}
		for _, d := range departments {
			DB.Create(&d)
		}
	}
}

func SeedStacks() {
	var count int64
	DB.Model(&model.Stack{}).Count(&count)
	if count == 0 {
		stacks := []model.Stack{
			{ID: 1, Name: "Default Stack"},
		}
		for _, s := range stacks {
			DB.Create(&s)
		}
	}
}

func SeedPositions() {
	var count int64
	DB.Model(&model.Position{}).Count(&count)
	if count == 0 {
		positions := []model.Position{
			{ID: 1, Name: "Default Position"},
		}
		for _, p := range positions {
			DB.Create(&p)
		}
	}
}
