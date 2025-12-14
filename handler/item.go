package handler

import (
	"github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/dto"
	"github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/model"
	"github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/service"
)

type HandlerItem struct {
	ServiceItem service.ServiceItemInterface
}

func NewHandlerItem(serviceItem service.ServiceItemInterface) HandlerItem {
	return HandlerItem{
		ServiceItem: serviceItem,
	}
}

func (handler *HandlerItem) GetItems(p dto.ItemParam) ([]model.Item, error) {
	return handler.ServiceItem.GetItems(p)
}

func (handler *HandlerItem) MustReplacedItems() ([]model.Item, error) {
	return handler.ServiceItem.MustReplacedItems()
}

func (handler *HandlerItem) GetItemByID(id int) (*model.Item, error) {
	return handler.ServiceItem.GetItemByID(id)
}

func (handler *HandlerItem) CreateItem(r dto.CreateItemRequest) (*model.Item, error) {
	return handler.ServiceItem.CreateItem(r)
}

func (handler *HandlerItem) UpdateItem(id int, r dto.UpdateItemRequest) error {
	return handler.ServiceItem.UpdateItem(id, r)
}

func (handler *HandlerItem) DeleteItem(id int) error {
	return handler.ServiceItem.DeleteItem(id)
}