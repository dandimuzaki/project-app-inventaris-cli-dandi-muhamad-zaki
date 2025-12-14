package service

import (
	"github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/repository"
)

type Service struct {
	ServiceCategory ServiceCategory
	ServiceItem     ServiceItem
	ServiceReport ServiceReport
}

func NewService(repo repository.Repository) Service {
	return Service{
		ServiceCategory: NewServiceCategory(&repo.RepoCategory),
		ServiceItem: NewServiceItem(&repo.RepoItem),
		ServiceReport: NewServiceReport(&repo.RepoReport),
	}
}