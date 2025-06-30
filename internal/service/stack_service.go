package service

import (
	"my-gin-app/internal/model"
	"my-gin-app/internal/repository"
)

func GetAllStacks() ([]model.Stack, error) {
	return repository.FindAllStacks()
}

func AddStack(stack *model.Stack) error {
	return repository.CreateStack(stack)
}
