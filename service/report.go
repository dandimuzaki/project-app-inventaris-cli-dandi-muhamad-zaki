package service

import (
	"session-14/model"
	"session-14/repository"
)

type ServiceReportInterface interface {
	GetReportMonthly(status string) ([]model.Report, error)
	GetLoyalCustomers(status string) ([]model.LoyalCustomer, error)
	GetBusyAreas() ([]model.BusyArea, error)
	GetBusyTimes() ([]model.BusyTime, error)
}

type ServiceReport struct {
	RepoReport repository.RepositoryReportInterface
}

func NewServiceReport(repoReport repository.RepositoryReportInterface) ServiceReport {
	return ServiceReport{
		RepoReport: repoReport,
	}
}

func (serviceReport *ServiceReport) GetReportMonthly(status string) ([]model.Report, error) {
	return serviceReport.RepoReport.GetReportMonthly(status)
}

func (serviceReport *ServiceReport) GetLoyalCustomers(status string) ([]model.LoyalCustomer, error) {
	return serviceReport.RepoReport.GetLoyalCustomers(status)
}

func (serviceReport *ServiceReport) GetBusyAreas() ([]model.BusyArea, error) {
	return serviceReport.RepoReport.GetBusyAreas()
}

func (serviceReport *ServiceReport) GetBusyTimes() ([]model.BusyTime, error) {
	return serviceReport.RepoReport.GetBusyTimes()
}
