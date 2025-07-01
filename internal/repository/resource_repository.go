package repository

import "my-gin-app/internal/model"

func FindResourceByName(name string) (model.Resource, error) {
	var resource model.Resource
	result := DB.Where("name = ?", name).First(&resource)
	return resource, result.Error
}
