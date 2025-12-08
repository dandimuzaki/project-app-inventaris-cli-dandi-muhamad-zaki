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
		fmt.Printf("Month : %s | Total_order : %d\n", t.Month, t.TotalOrder)
	}
}

func (handlerReport *HandlerReport) MonthlyLoyalCustomers(status string) {
	// call service report monthly
	customers, err := handlerReport.ServiceReport.GetLoyalCustomers(status)
	if err != nil {
		fmt.Println("Failed to retrieve data:", err)
	}

	fmt.Println("Loyal Customers per Month")
	for _, t := range customers {
		fmt.Printf("Month : %s | Customer Name: %s | Total_order : %d\n", t.Month, t.CustomerName, t.TotalOrder)
	}
}

func (handlerReport *HandlerReport) BusyAreas() {
	// call service report monthly
	areas, err := handlerReport.ServiceReport.GetBusyAreas()
	if err != nil {
		fmt.Println("Failed to retrieve data:", err)
	}

	fmt.Println("Busy Areas")
	for _, t := range areas {
		fmt.Printf("Area : %s | Total_order : %d\n", t.Area, t.TotalOrder)
	}
}

func (handlerReport *HandlerReport) BusyTimes() {
	// call service report monthly
	times, err := handlerReport.ServiceReport.GetBusyTimes()
	if err != nil {
		fmt.Println("Failed to retrieve data:", err)
	}

	fmt.Println("Busy Areas")
	for _, t := range times {
		fmt.Printf("Hour : %s | Total_order : %d\n", t.Hour, t.TotalOrder)
	}
}
