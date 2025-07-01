package service

import (
	"my-gin-app/internal/model"
	"my-gin-app/internal/repository"
)

func GetAllFeatureFlags() ([]model.FeatureFlag, error) {
	return repository.FindAllFeatureFlags()
}

func AddFeatureFlag(flag *model.FeatureFlag) error {
	return repository.CreateFeatureFlag(flag)
}
