package app

import (
	"account/app/models"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (app *App) CreateUser(c *gin.Context) {
	var newUser models.GetCreateAccount

	if err := c.ShouldBindJSON(&newUser); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	newUser.Sub = uuid.New().String()
	_, err := newUser.Create(app.Collection)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	Publish("register", fmt.Sprintf("New User registered with %s id", newUser.Sub), app.MQChannel, app.MQQueue)
	c.JSON(http.StatusCreated, gin.H{"success": true})
}

func (app *App) GetUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		log.Println("Empty account id")
	}

	var account models.GetCreateAccount
	account.Sub = id
	err := account.Get(app.Collection)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "User does not exists"})
		return
	}

	c.JSON(http.StatusOK, account)
}

func (app *App) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		log.Println("Empty account id")
	}

	var payload models.UpdateAccount
	if err := c.ShouldBindJSON(&payload); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if reflect.ValueOf(payload).IsZero() {
		log.Println(fmt.Errorf("payload %v is empty or contain wrong values. Nothing to udpate", payload))
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Cant handle payload (empty or incorrect)"})
		return
	}

	err := models.Update(app.Collection, id, payload)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to update account"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}
