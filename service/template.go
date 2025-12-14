package service

import (
	"github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/model"
	"github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/repository"
)

type ServiceTemplate struct {
	RepoTemplate repository.RepositoryTemplate
}

func NewServiceTemplate(repoTemplate repository.RepositoryTemplate) ServiceTemplate {
	return ServiceTemplate{
		RepoTemplate: repoTemplate,
	}
}

func (serviceTemplate *ServiceTemplate) Create(template *model.Template) error {
	return serviceTemplate.RepoTemplate.Create(template)
}

func (serviceTemplate *ServiceTemplate) List() ([]*model.Template, error) {
	return serviceTemplate.RepoTemplate.List()
}

func (serviceTemplate *ServiceTemplate) Update(id int, template *model.Template) error {
	return serviceTemplate.RepoTemplate.Update(id, template)
}

func (serviceTemplate *ServiceTemplate) Delete(id int) error {
	return serviceTemplate.RepoTemplate.Delete(id)
}
