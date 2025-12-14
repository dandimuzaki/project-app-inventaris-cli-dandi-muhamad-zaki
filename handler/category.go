package handler

import (
	"github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/dto"
	"github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/model"
	"github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/service"
)

type HandlerCategory struct {
	ServiceCategory service.ServiceCategoryInterface
}

func NewHandlerCategory(serviceCategory service.ServiceCategoryInterface) HandlerCategory {
	return HandlerCategory{
		ServiceCategory: serviceCategory,
	}
}

func (handlerCategory *HandlerCategory) GetCategories() ([]model.Category, error) {
	return handlerCategory.ServiceCategory.GetCategories()
}

func (handlerCategory *HandlerCategory) GetCategoryByID(id int) (*model.Category, error) {
	return handlerCategory.ServiceCategory.GetCategoriesByID(id)
}

func (handlerCategory *HandlerCategory) CreateCategory(r dto.CreateCategoryRequest) (*model.Category, error) {
	return handlerCategory.ServiceCategory.CreateCategory(r)
}

func (handlerCategory *HandlerCategory) UpdateCategory(id int, r dto.UpdateCategoryRequest) error {
	return handlerCategory.ServiceCategory.UpdateCategory(id, r)
}

func (handlerCategory *HandlerCategory) DeleteCategory(id int) error {
	return handlerCategory.ServiceCategory.DeleteCategory(id)
}