package database

import "products/app/models"

type QueryParams struct {
	Page    int64
	Offset  int64
	PerPage int64
}

type ProductRepository interface {
	GetAll() []models.Product
	Get(product_id int64) (*models.Product, error)
	Create(product *models.Product) error
	Delete(product_id int64) error
	Patch(product_id int64, payload *models.Product) error
}
