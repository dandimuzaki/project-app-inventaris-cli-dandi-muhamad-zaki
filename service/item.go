package service

import (
	"github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/dto"
	"github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/model"
	"github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/repository"
)

type ServiceItemInterface interface {
	GetItems(p dto.ItemParam) ([]model.Item, error)
	MustReplacedItems() ([]model.Item, error)
	GetItemByID(id int) (*model.Item, error)
	CreateItem(r dto.CreateItemRequest) (*model.Item, error)
	UpdateItem(id int, r dto.UpdateItemRequest) error
	DeleteItem(id int) error
}

type ServiceItem struct {
	RepoItem repository.RepositoryItemInterface
}

func NewServiceItem(repo repository.RepositoryItemInterface) ServiceItem {
	return ServiceItem{
		RepoItem: repo,
	}
}

func (service *ServiceItem) GetItems(p dto.ItemParam) ([]model.Item, error) {
	return service.RepoItem.GetItems(p)
}

func (service *ServiceItem) MustReplacedItems() ([]model.Item, error) {
	return service.RepoItem.MustReplacedItems()
}

func (service *ServiceItem) GetItemByID(id int) (*model.Item, error) {
	return service.RepoItem.GetItemByID(id)
}

func (service *ServiceItem) CreateItem(r dto.CreateItemRequest) (*model.Item, error) {
	return service.RepoItem.CreateItem(r)
}

func (service *ServiceItem) UpdateItem(id int, r dto.UpdateItemRequest) error {
	return service.RepoItem.UpdateItem(id, r)
}

func (service *ServiceItem) DeleteItem(id int) error {
	return service.RepoItem.DeleteItem(id)
}