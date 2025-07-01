package repository

import "my-gin-app/internal/model"

func FindResourceByName(name string) (model.Resource, error) {
	var resource model.Resource
	result := DB.Where("name = ?", name).First(&resource)
	return resource, result.Error
}

func FindAllResources() ([]model.Resource, error) {
	var resources []model.Resource
	result := DB.Find(&resources)
	return resources, result.Error
}

func CreateResource(resource *model.Resource) error {
	return DB.Create(resource).Error
}
