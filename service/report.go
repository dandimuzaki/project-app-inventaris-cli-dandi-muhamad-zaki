package service

import "github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/repository"

type ServiceReportInterface interface {
	TotalNetValue() (*float64, error)
	TotalInvestment() (*float64, error)
}

type ServiceReport struct {
	RepoReport repository.RepositoryReportInterface
}

func NewServiceReport(repoReport repository.RepositoryReportInterface) ServiceReport {
	return ServiceReport{
		RepoReport: repoReport,
	}
}

func (serviceReport *ServiceReport) TotalNetValue() (*float64, error) {
	return serviceReport.RepoReport.TotalNetValue()
}

func (serviceReport *ServiceReport) TotalInvestment() (*float64, error) {
	return serviceReport.RepoReport.TotalInvestment()
}