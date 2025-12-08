package repository

import (
	"context"
	"session-14/database"
	"session-14/model"
)

type RepositoryReportInterface interface {
	GetReportMonthly(status string) ([]model.Report, error)
	GetLoyalCustomers(status string) ([]model.LoyalCustomer, error)
	GetBusyAreas() ([]model.BusyArea, error)
	GetBusyTimes() ([]model.BusyTime, error)
}

type RepositoryReport struct {
	DB database.PgxIface
}

func NewRepoReport(db database.PgxIface) RepositoryReport {
	return RepositoryReport{
		DB: db,
	}
}

func (repo *RepositoryReport) GetReportMonthly(status string) ([]model.Report, error) {
	query := `SELECT TO_CHAR(DATE_TRUNC('month', datetime_order), 'Month') AS month, COUNT(*) AS total_order
FROM orders
WHERE status = $1
GROUP BY month
ORDER BY MIN(DATE_TRUNC('month', datetime_order))`

	rows, err := repo.DB.Query(context.Background(), query, status)
	if err != nil {
		return nil, err
	}

	var reportMonthly []model.Report
	var report model.Report
	for rows.Next() {
		err := rows.Scan(&report.Month, &report.TotalOrder)
		if err != nil {
			return nil, err
		}
		reportMonthly = append(reportMonthly, report)
	}

	return reportMonthly, nil
}

func (repo *RepositoryReport) GetLoyalCustomers(status string) ([]model.LoyalCustomer, error) {
	query := `WITH monthly_orders AS (
    SELECT
        DATE_TRUNC('month', r.start_time) AS bulan,
        c.id AS customer_id,
        c.name AS customer_name,
        COUNT(CASE WHEN r.status = $1 THEN r.id END) AS total_order,
        RANK() OVER (
            PARTITION BY DATE_TRUNC('month', r.start_time)
            ORDER BY COUNT(CASE WHEN r.status = $1 THEN r.id END) DESC
        ) AS rank_in_month
    FROM rides r
	LEFT JOIN ride_requests rr
	ON r.request_id = rr.id
    LEFT JOIN customers c ON rr.customer_id = c.id
    GROUP BY bulan, c.id, c.name
)
SELECT
    TO_CHAR(bulan, 'Month') AS "month",
    customer_name,
    total_order
FROM monthly_orders
WHERE rank_in_month = 1
ORDER BY bulan`

	rows, err := repo.DB.Query(context.Background(), query, status)
	if err != nil {
		return nil, err
	}

	var loyalCustomers []model.LoyalCustomer
	var customer model.LoyalCustomer
	for rows.Next() {
		err := rows.Scan(&customer.Month, &customer.CustomerName, &customer.TotalOrder)
		if err != nil {
			return nil, err
		}
		loyalCustomers = append(loyalCustomers, customer)
	}

	return loyalCustomers, nil
}

func (repo *RepositoryReport) GetBusyAreas() ([]model.BusyArea, error) {
	query := `SELECT
	d.name AS area,
	COUNT(rr.id) AS total_order
FROM ride_requests rr
LEFT JOIN locations AS l
ON rr.pickup_location_id = l.id
LEFT JOIN subdistricts AS sd
ON l.subdistrict_id = sd.id
LEFT JOIN districts AS d
ON sd.district_id = d.id
GROUP BY d.name
ORDER BY total_order DESC`

	rows, err := repo.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	var busyAreas []model.BusyArea
	var area model.BusyArea
	for rows.Next() {
		err := rows.Scan(&area.Area, &area.TotalOrder)
		if err != nil {
			return nil, err
		}
		busyAreas = append(busyAreas, area)
	}

	return busyAreas, nil
}

func (repo *RepositoryReport) GetBusyTimes() ([]model.BusyTime, error) {
	query := `SELECT
	TO_CHAR(DATE_TRUNC('hour', created_at), 'HH') AS "hour",
	COUNT(id) AS "total_order"
FROM ride_requests
GROUP BY "hour"
ORDER BY "total_order" DESC`

	rows, err := repo.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	var busyTimes []model.BusyTime
	var time model.BusyTime
	for rows.Next() {
		err := rows.Scan(&time.Hour, &time.TotalOrder)
		if err != nil {
			return nil, err
		}
		busyTimes = append(busyTimes, time)
	}

	return busyTimes, nil
}