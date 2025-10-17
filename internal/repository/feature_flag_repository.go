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

func FindFeatureFlagByID(id int) (model.FeatureFlag, error) {
	var flag model.FeatureFlag
	result := DB.Where("id = ?", id).First(&flag)
	return flag, result.Error
}

func UpdateFeatureFlag(flag *model.FeatureFlag) error {
	return DB.Save(flag).Error
}

func DeleteFeatureFlag(id int) error {
	return DB.Delete(&model.FeatureFlag{}, id).Error
}
