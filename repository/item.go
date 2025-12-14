package repository

import (
	"context"

	"github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/database"
	"github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/dto"
	"github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/model"
	"github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/utils"
	"github.com/jackc/pgx/v5"
)

// Initialize constant
var depreciationRate float64 = 0.2
var limitUsage int = 100

type RepositoryItemInterface interface {
	GetItems(p dto.ItemParam) ([]model.Item, error)
	MustReplacedItems() ([]model.Item, error)
	GetItemByID(id int) (*model.Item, error)
	CreateItem(r dto.CreateItemRequest) (*model.Item, error)
	UpdateItem(id int, r dto.UpdateItemRequest) error
	DeleteItem(id int) error
}

type RepositoryItem struct {
	DB database.PgxIface
}

func NewRepoItem(db database.PgxIface) RepositoryItem {
	return RepositoryItem{
		DB: db,
	}
}

func (repo *RepositoryItem) GetItems(p dto.ItemParam) ([]model.Item, error) {
	query := `SELECT
		i.id, i.item_name,
		c.category_name, i.price,
		i.purchase_date,
		EXTRACT(DAY FROM(NOW() - purchase_date)) AS total_usage,
		FLOOR(EXTRACT(DAY FROM(NOW() - purchase_date))/365) AS years_in_service,
		(1.0-$1)^(FLOOR(EXTRACT(DAY FROM(NOW() - purchase_date))/365)) * price AS net_value
	FROM items i
	LEFT JOIN categories c
	ON i.category_id = c.id
	WHERE LOWER(i.item_name) LIKE '%'||$2||'%'
	AND i.category_id = COALESCE($3, i.category_id)
	ORDER BY id`

	rows, err := repo.DB.Query(context.Background(), query, depreciationRate, p.Keyword, p.CategoryID)
	if err != nil {
		return nil, err
	}

	var items []model.Item
	var item model.Item
	for rows.Next() {
		err := rows.Scan(&item.ID, &item.Name, &item.Category, &item.Price, &item.PurchaseDate, &item.TotalUsage, &item.YearsInService, &item.NetValue)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (repo *RepositoryItem) MustReplacedItems() ([]model.Item, error) {
	query := `SELECT
		i.id, i.item_name,
		c.category_name, i.price,
		i.purchase_date,
		EXTRACT(DAY FROM(NOW() - purchase_date)) AS total_usage,
		FLOOR(EXTRACT(DAY FROM(NOW() - purchase_date))/365) AS years_in_service,
		(1.0-$1)^FLOOR(EXTRACT(DAY FROM(NOW() - purchase_date))/365) * price AS net_value
	FROM items i
	LEFT JOIN categories c
	ON i.category_id = c.id
	WHERE EXTRACT(DAY FROM(NOW() - purchase_date)) > $2
	ORDER BY EXTRACT(DAY FROM(NOW() - purchase_date)) DESC`

	rows, err := repo.DB.Query(context.Background(), query, depreciationRate, limitUsage)
	if err != nil {
		return nil, err
	}

	var items []model.Item
	var item model.Item
	for rows.Next() {
		err := rows.Scan(&item.ID, &item.Name, &item.Category, &item.Price, &item.PurchaseDate, &item.TotalUsage, &item.YearsInService, &item.NetValue)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (repo *RepositoryItem) GetItemByID(id int) (*model.Item, error) {
	query := `SELECT
		i.id, i.item_name,
		c.category_name, i.price,
		i.purchase_date,
		EXTRACT(DAY FROM(NOW() - purchase_date)) AS total_usage,
		FLOOR(EXTRACT(DAY FROM(NOW() - purchase_date))/365) AS years_in_service,
		(1.0-$1)^FLOOR(EXTRACT(DAY FROM(NOW() - purchase_date))/365) * price AS net_value
	FROM items i
	LEFT JOIN categories c
	ON i.category_id = c.id
	WHERE i.id = $2`

	var item model.Item
	err := repo.DB.QueryRow(context.Background(), query, depreciationRate, id).Scan(&item.ID, &item.Name, &item.Category, &item.Price, &item.PurchaseDate, &item.TotalUsage, &item.YearsInService, &item.NetValue)

	if err == pgx.ErrNoRows {
		return nil, utils.ErrItemNotFound
	} else if err != nil && err != pgx.ErrNoRows {
		return nil, err
	}

	return &item, nil
}

func (repo *RepositoryItem) CreateItem(r dto.CreateItemRequest) (*model.Item, error) {
	query := `INSERT INTO items (item_name, category_id, price, purchase_date)
VALUES ($1, $2, $3, $4)
RETURNING 
  id,
  item_name,
  (SELECT category_name FROM categories WHERE id = $2) AS category_name,
  price,
  purchase_date,
	EXTRACT(DAY FROM(NOW() - $4)) AS total_usage,
	FLOOR(EXTRACT(DAY FROM(NOW() - $4))/365) AS years_in_service,
	(1.0-$5)^FLOOR(EXTRACT(DAY FROM(NOW() - $4))/365) * $3 AS net_value;
`

	var item model.Item
	err := repo.DB.QueryRow(context.Background(), query, r.Name, r.CategoryId, r.Price, r.PurchaseDate, depreciationRate).Scan(&item.ID, &item.Name, &item.Category, &item.Price, &item.PurchaseDate, &item.TotalUsage, &item.YearsInService, &item.NetValue)

	if err == pgx.ErrNoRows {
		return nil, utils.ErrItemNotFound
	} else if err != nil && err != pgx.ErrNoRows {
		return nil, err
	}

	return &item, nil
}

func (repo *RepositoryItem) UpdateItem(id int, r dto.UpdateItemRequest) error {
	query := `UPDATE items SET
		item_name = COALESCE($1, item_name),
		category_id = COALESCE($2, category_id),
		price = COALESCE($3, price),
		purchase_date = COALESCE($4, purchase_date)
		WHERE id = $5
	`
	commandTag, err := repo.DB.Exec(context.Background(), query, r.Name, r.CategoryId, r.Price, r.PurchaseDate, id)
	if commandTag.RowsAffected() != 1 {
		return utils.ErrItemNotFound
	}
	return err
}

func (repo *RepositoryItem) DeleteItem(id int) error {
	query := `DELETE FROM items
		WHERE id = $1
	`
	commandTag, err := repo.DB.Exec(context.Background(), query, id)
	if commandTag.RowsAffected() != 1 {
		return utils.ErrItemNotFound
	}
	return err
}