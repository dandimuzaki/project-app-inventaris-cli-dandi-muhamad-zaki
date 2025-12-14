package repository

import (
	"context"

	"github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/database"
)

type RepositoryReportInterface interface {
	TotalNetValue() (*float64, error)
	TotalInvestment() (*float64, error)
}

type RepositoryReport struct {
	DB database.PgxIface
}

func NewRepoReport(db database.PgxIface) RepositoryReport {
	return RepositoryReport{
		DB: db,
	}
}

func (repo *RepositoryReport) TotalNetValue() (*float64, error) {
	query := `SELECT COALESCE(SUM((1.0-$1)^FLOOR(EXTRACT(DAY FROM(NOW() - purchase_date))/365) * price), 0) AS total_net_value
	FROM items`

	var totalNetValue float64
	err := repo.DB.QueryRow(context.Background(), query, depreciationRate).Scan(&totalNetValue)
	if err != nil {
		return nil, err
	}

	return &totalNetValue, nil
}

func (repo *RepositoryReport) TotalInvestment() (*float64, error) {
	query := `SELECT COALESCE(SUM(price), 0) AS total_investment
	FROM items`

	var totalInvestment float64
	err := repo.DB.QueryRow(context.Background(), query).Scan(&totalInvestment)
	if err != nil {
		return nil, err
	}

	return &totalInvestment, nil
}