package dto

import (
	"database/sql"
	"time"
)

type CreateCategoryRequest struct {
	Name string 
	Description sql.NullString
}

type UpdateCategoryRequest struct {
	Name sql.NullString
	Description sql.NullString
}

type CreateItemRequest struct {
	Name                 string
	CategoryId int
	Price                float64
	PurchaseDate         time.Time
}

type UpdateItemRequest struct {
	Name                 sql.NullString
	CategoryId sql.NullInt32
	Price                sql.NullFloat64
	PurchaseDate         sql.NullTime
}

type ItemParam struct {
	Keyword string
	CategoryID sql.NullInt32
}