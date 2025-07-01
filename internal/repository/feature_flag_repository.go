package repository

import "my-gin-app/internal/model"

func FindAllFeatureFlags() ([]model.FeatureFlag, error) {
	var flags []model.FeatureFlag
	result := DB.Find(&flags)
	return flags, result.Error
}

func CreateFeatureFlag(flag *model.FeatureFlag) error {
	return DB.Create(flag).Error
}
