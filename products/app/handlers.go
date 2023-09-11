package app

import (
	"net/http"
	"products/app/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (app *App) ProductsList(c *gin.Context) {
	// TODO: filter and pagination by the query params
}

func (app *App) ProductCreate(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err})
		return
	}
	if err := app.controller.Save(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (app *App) ProductDetail(c *gin.Context) {
	productId, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product id"})
		return
	}
	product, err := app.controller.Get(productId)
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	c.JSON(http.StatusOK, product)
}

func (app *App) ProductPatch(c *gin.Context) {
	productId, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product id"})
		return
	}
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err})
		return
	}

	err = app.controller.Patch(productId, &product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (app *App) ProductDelete(c *gin.Context) {
	productId, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product id"})
		return
	}
	go app.controller.Delete(productId)
	c.JSON(http.StatusOK, nil)
}
