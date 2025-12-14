package handler

import "github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/service"

type TemplateHandler struct {
	Service service.ServiceTemplate
}

func NewTemplateHandler(service service.ServiceTemplate) *TemplateHandler {
	return &TemplateHandler{
		Service: service,
	}
}
