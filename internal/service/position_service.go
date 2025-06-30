package service

import (
	"my-gin-app/internal/model"
	"my-gin-app/internal/repository"
)

func GetAllPositions() ([]model.Position, error) {
	return repository.FindAllPositions()
}

func AddPosition(pos *model.Position) error {
	return repository.CreatePosition(pos)
}
