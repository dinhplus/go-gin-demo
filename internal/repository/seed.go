package repository

import "my-gin-app/internal/model"

func SeedDepartments() {
	var count int64
	DB.Model(&model.Department{}).Count(&count)
	if count == 0 {
		departments := []model.Department{
			{Name: "Engineering"},
			{Name: "HR"},
			{Name: "Sales"},
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
			{Name: "Backend"},
			{Name: "Frontend"},
			{Name: "DevOps"},
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
			{Name: "Manager"},
			{Name: "Developer"},
			{Name: "Intern"},
		}
		for _, p := range positions {
			DB.Create(&p)
		}
	}
}
