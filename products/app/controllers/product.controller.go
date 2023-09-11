package controllers

import (
	"database/sql"
	"products/app/models"
)

type ProductController struct {
	DB *sql.DB
}

func Init(postgres_url string) (*ProductController, error) {
	pc := &ProductController{}
	return pc, nil
}

func (pc *ProductController) GetAll() {
	// page, offset, perpage int64
}

func (pc *ProductController) SearchByName(q string) {
	// , page, offset, perpage int64
}

func (pc *ProductController) Get(id int64) (*models.Product, error) {

	return nil, nil
}

func (pc *ProductController) Save(p *models.Product) error {
	p.Init()

	return nil
}

func (pc *ProductController) Delete(id int64) {

}

func (pc *ProductController) Patch(id int64, payload *models.Product) error {
	return nil
}
