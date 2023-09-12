package database

import (
	"database/sql"
	"log"
	"products/app/models"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/lib/pq"
)

type QueryParams struct {
	Page    int64
	Offset  int64
	PerPage int64
}

type ProductRepository interface {
	GetAll() ([]models.Product, error)
	Get(product_id int64) (*models.Product, error)
	Create(product *models.Product) error
	Delete(product_id int64) error
	Patch(product_id int64, payload *models.Product) error
}

type ProductDBController struct {
	db *sql.DB
}

func Init(postgres_url string) (*ProductDBController, error) {
	db, err := sql.Open("postgres", postgres_url)
	if err != nil {
		return nil, err
	}
	pc := &ProductDBController{db: db}
	return pc, nil
}

func (pdbc *ProductDBController) GetAll() ([]models.Product, error) {
	q := goqu.From("products")
	selectQuery, _, _ := q.ToSQL()
	rows, err := pdbc.db.Query(selectQuery)
	if err != nil {
		return nil, err
	}

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.ProductCode,
			&product.Price,
			&product.CreatedAt,
		); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func (pdbc *ProductDBController) Get(product_id int64) (*models.Product, error) {
	q := goqu.From("products").Where(goqu.C("id").Eq(product_id))
	selectQuery, _, _ := q.ToSQL()
	var product models.Product
	if err := pdbc.db.QueryRow(selectQuery).Scan(
		&product.Id,
		&product.Name,
		&product.ProductCode,
		&product.Price,
		&product.CreatedAt,
	); err != nil {
		return nil, err
	}
	return &product, nil
}

func (pdbc *ProductDBController) Create(product *models.Product) error {
	q := goqu.Insert("products").Rows(
		goqu.Record{
			"name":         product.Name,
			"price":        product.Price,
			"product_code": product.ProductCode,
		},
	).Returning(
		"id", "product_code", "created_at",
	)
	insertSQL, _, _ := q.ToSQL()
	err := pdbc.db.QueryRow(insertSQL).Scan(
		&product.Id, &product.ProductCode, &product.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (pdbc *ProductDBController) Delete(product_id int64) error {
	q := goqu.Delete("products").Where(goqu.C("id").Eq(product_id))
	deleteSQL, _, _ := q.ToSQL()
	res, err := pdbc.db.Exec(deleteSQL)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	log.Println("Rows deleted: ", rows)
	return nil
}

func (pdbc *ProductDBController) Patch(product_id int64, payload *models.Product) error {
	q := goqu.Update("products").Set(
		goqu.Record{"name": payload.Name, "price": payload.Price},
	).Where(goqu.C("id").Eq(product_id))
	updateSQL, _, _ := q.ToSQL()
	res, err := pdbc.db.Exec(updateSQL)
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	log.Println("Rows updated: ", rows)
	return nil
}
