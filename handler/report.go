package handler

import "github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/service"

type HandlerReport struct {
	ServiceReport service.ServiceReportInterface
}

func NewHandlerReport(serviceReport service.ServiceReportInterface) HandlerReport {
	return HandlerReport{
		ServiceReport: serviceReport,
	}
}

func (handlerReport *HandlerReport) TotalNetValue() (*float64, error) {
	return handlerReport.ServiceReport.TotalNetValue()
}

func (handlerReport *HandlerReport) TotalInvestment() (*float64, error) {
	return handlerReport.ServiceReport.TotalInvestment()
}