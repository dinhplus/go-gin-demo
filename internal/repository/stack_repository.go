package repository

import "my-gin-app/internal/model"

func FindAllStacks() ([]model.Stack, error) {
	var stacks []model.Stack
	result := DB.Find(&stacks)
	return stacks, result.Error
}

func CreateStack(stack *model.Stack) error {
	return DB.Create(stack).Error
}
