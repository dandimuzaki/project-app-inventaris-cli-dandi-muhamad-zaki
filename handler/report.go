package handler

import (
	"fmt"
	"session-14/service"
)

type HandlerReport struct {
	ServiceReport service.ServiceReportInterface
}

func NewHandlerReport(serviceReport service.ServiceReportInterface) HandlerReport {
	return HandlerReport{
		ServiceReport: serviceReport,
	}
}

func (handlerReport *HandlerReport) ReportMonthly(status string) {
	// call service report monthly
	reports, err := handlerReport.ServiceReport.GetReportMonthly(status)
	if err != nil {
		fmt.Println("Failed to retrieve data:", err)
	}

	fmt.Println("Order per Month")
	for _, t := range reports {
		fmt.Printf("Month : %s, Total_order : %d\n", t.Month, t.TotalOrder)
	}
}
