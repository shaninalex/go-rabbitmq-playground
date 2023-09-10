package controllers

import (
	"database/sql"
	"products/app/models"
)

type ProductController struct {
	DB *sql.DB
}

func (pc *ProductController) Get(id int64)                                       {}
func (pc *ProductController) Save(p *models.Product)                             {}
func (pc *ProductController) Delete(id int64)                                    {}
func (pc *ProductController) GetAll(page, offset, perpage int64)                 {}
func (pc *ProductController) SearchByName(q string, page, offset, perpage int64) {}
