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

func FindStackByID(id int) (model.Stack, error) {
	var stack model.Stack
	result := DB.Where("id = ?", id).First(&stack)
	return stack, result.Error
}

func UpdateStack(stack *model.Stack) error {
	return DB.Save(stack).Error
}

func DeleteStack(id int) error {
	return DB.Delete(&model.Stack{}, id).Error
}
