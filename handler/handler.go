package handler

import "github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/service"

type Handler struct {
	HandlerCategory HandlerCategory
	HandlerItem     HandlerItem
	HandlerReport HandlerReport
}

func NewHandler(service service.Service) Handler {
	return Handler{
		HandlerCategory: NewHandlerCategory(&service.ServiceCategory),
		HandlerItem: NewHandlerItem(&service.ServiceItem),
		HandlerReport: NewHandlerReport(&service.ServiceReport),
	}
}