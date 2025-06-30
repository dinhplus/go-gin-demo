package repository

import "my-gin-app/internal/model"

func FindAllPositions() ([]model.Position, error) {
	var positions []model.Position
	result := DB.Find(&positions)
	return positions, result.Error
}

func CreatePosition(pos *model.Position) error {
	return DB.Create(pos).Error
}
