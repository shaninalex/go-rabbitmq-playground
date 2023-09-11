package controllers

import (
	"products/app/database"
	"products/app/models"
)

type ProductController struct {
	DB database.ProductRepository
}

func Init(postgres_url string) (*ProductController, error) {
	pc := &ProductController{}
	return pc, nil
}

func (pc *ProductController) GetAll() []models.Product {
	// TODO: handle pagination (page, offset, perpage)
	products := pc.DB.GetAll()
	return products
}

func (pc *ProductController) Get(id int64) (*models.Product, error) {
	product, err := pc.DB.Get(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (pc *ProductController) Save(p *models.Product) error {
	p.Init()
	if err := pc.DB.Create(p); err != nil {
		return err
	}
	return nil
}

func (pc *ProductController) Delete(id int64) error {
	if err := pc.DB.Delete(id); err != nil {
		return err
	}
	return nil

}

func (pc *ProductController) Patch(id int64, payload *models.Product) error {
	if err := pc.DB.Patch(id, payload); err != nil {
		return err
	}
	return nil
}
