package app

import (
	"account/app/models"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (app *App) CreateUser(c *gin.Context) {
	var newUser models.Account

	if err := c.BindJSON(&newUser); err != nil {
		// DO SOMETHING WITH THE ERROR
		log.Println(err)
	}

	result, err := app.Collection.InsertOne(c, newUser)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(result)
}

func (app *App) GetUser(c *gin.Context) {

}

func (app *App) UpdateUser(c *gin.Context) {

}
