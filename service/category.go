package service

import (
	"github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/dto"
	"github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/model"
	"github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/repository"
)

type ServiceCategoryInterface interface {
	GetCategories() ([]model.Category, error)
	GetCategoriesByID(id int) (*model.Category, error)
	CreateCategory(r dto.CreateCategoryRequest) (*model.Category, error)
	UpdateCategory(id int, r dto.UpdateCategoryRequest) error
	DeleteCategory(id int) error
}

type ServiceCategory struct {
	RepoCategory repository.RepositoryCategoryInterface
}

func NewServiceCategory(repoCategory repository.RepositoryCategoryInterface) ServiceCategory {
	return ServiceCategory{
		RepoCategory: repoCategory,
	}
}

func (serviceCategory *ServiceCategory) GetCategories() ([]model.Category, error) {
	return serviceCategory.RepoCategory.GetCategories()
}

func (serviceCategory *ServiceCategory) GetCategoriesByID(id int) (*model.Category, error) {
	return serviceCategory.RepoCategory.GetCategoryByID(id)
}

func (serviceCategory *ServiceCategory) CreateCategory(r dto.CreateCategoryRequest) (*model.Category, error) {
	return serviceCategory.RepoCategory.CreateCategory(r)
}

func (serviceCategory *ServiceCategory) UpdateCategory(id int, r dto.UpdateCategoryRequest) error {
	return serviceCategory.RepoCategory.UpdateCategory(id, r)
}

func (serviceCategory *ServiceCategory) DeleteCategory(id int) error {
	return serviceCategory.RepoCategory.DeleteCategory(id)
}