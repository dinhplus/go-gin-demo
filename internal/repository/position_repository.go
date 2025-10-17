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

func FindPositionByID(id int) (model.Position, error) {
	var pos model.Position
	result := DB.Where("id = ?", id).First(&pos)
	return pos, result.Error
}

func UpdatePosition(pos *model.Position) error {
	return DB.Save(pos).Error
}

func DeletePosition(id int) error {
	return DB.Delete(&model.Position{}, id).Error
}
