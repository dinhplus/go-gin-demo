package service

import (
	"my-gin-app/internal/model"
	"my-gin-app/internal/repository"
)

func GetAllResources() ([]model.Resource, error) {
	return repository.FindAllResources()
}

func AddResource(resource *model.Resource) error {
	return repository.CreateResource(resource)
}
