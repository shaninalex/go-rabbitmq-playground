package controllers

import (
	"products/app/database"
	"products/app/models"
)

type ProductController struct {
	DB database.ProductRepository
}

func Init(postgres_url string) (*ProductController, error) {
	productDb, err := database.Init(postgres_url)
	if err != nil {
		return nil, err
	}
	pc := &ProductController{DB: productDb}
	return pc, nil
}

func (pc *ProductController) GetAll() ([]models.Product, error) {
	// TODO: handle pagination (page, offset, perpage)
	products, err := pc.DB.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (pc *ProductController) Get(id int64) (*models.Product, error) {
	product, err := pc.DB.Get(id)
	if err != nil {
		return nil, err
	}
	// TODO: Other actions
	return product, nil
}

func (pc *ProductController) Save(p *models.Product) error {
	p.Init()
	if err := pc.DB.Create(p); err != nil {
		return err
	}
	// TODO: Other actions
	return nil
}

func (pc *ProductController) Delete(id int64) error {
	if err := pc.DB.Delete(id); err != nil {
		return err
	}
	// TODO: Other actions
	return nil

}

func (pc *ProductController) Patch(id int64, payload *models.Product) error {
	if err := pc.DB.Patch(id, payload); err != nil {
		return err
	}
	// TODO: Other actions
	return nil
}
