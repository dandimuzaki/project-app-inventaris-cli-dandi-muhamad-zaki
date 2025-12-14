package repository

import (
	"context"

	"github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/database"
	"github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/dto"
	"github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/model"
	"github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/utils"
	"github.com/jackc/pgx/v5"
)

type RepositoryCategoryInterface interface {
	GetCategories() ([]model.Category, error)
	GetCategoryByID(id int) (*model.Category, error)
	CreateCategory(r dto.CreateCategoryRequest) (*model.Category, error)
	UpdateCategory(id int, r dto.UpdateCategoryRequest) error
	DeleteCategory(id int) error
}

type RepositoryCategory struct {
	DB database.PgxIface
}

func NewRepoCategory(db database.PgxIface) RepositoryCategory {
	return RepositoryCategory{
		DB: db,
	}
}

func (repo *RepositoryCategory) GetCategories() ([]model.Category, error) {
	query := `SELECT * FROM categories ORDER BY id`

	rows, err := repo.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	var categories []model.Category
	var category model.Category
	for rows.Next() {
		err := rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (repo *RepositoryCategory) GetCategoryByID(id int) (*model.Category, error) {
	query := `SELECT * FROM categories c WHERE c.id = $1`

	var category model.Category

	err := repo.DB.QueryRow(context.Background(), query, id).Scan(&category.ID, &category.Name, &category.Description)
	
	if err == pgx.ErrNoRows {
		return nil, utils.ErrCategoryNotFound
	} else if err != nil && err != pgx.ErrNoRows {
		return nil, err
	}

	return &category, nil
}

func (repo *RepositoryCategory) CreateCategory(r dto.CreateCategoryRequest) (*model.Category, error) {
	query := `INSERT INTO categories (category_name, description) VALUES ($1, $2) RETURNING id, category_name, description`

	var category model.Category

	err := repo.DB.QueryRow(context.Background(), query, r.Name, r.Description).Scan(&category.ID, &category.Name, &category.Description)
	
	if err == pgx.ErrNoRows {
		return nil, utils.ErrCategoryNotFound
	} else if err != nil && err != pgx.ErrNoRows {
		return nil, err
	}

	return &category, nil
}

func (repo *RepositoryCategory) UpdateCategory(id int, r dto.UpdateCategoryRequest) error {
	query := `UPDATE categories SET
		category_name = COALESCE($1, category_name),
		description = COALESCE($2, description)
		WHERE id = $3
	`
	commandTag, err := repo.DB.Exec(context.Background(), query, r.Name, r.Description, id)
	if commandTag.RowsAffected() != 1 {
		return utils.ErrCategoryNotFound
	}
	return err
}

func (repo *RepositoryCategory) DeleteCategory(id int) error {
	query := `DELETE FROM categories
		WHERE id = $1
	`
	commandTag, err := repo.DB.Exec(context.Background(), query, id)
	if commandTag.RowsAffected() != 1 {
		return utils.ErrCategoryNotFound
	}
	return err
}