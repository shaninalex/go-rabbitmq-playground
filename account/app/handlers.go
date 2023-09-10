package app

import (
	"account/app/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (app *App) CreateUser(c *gin.Context) {
	var newUser models.NewUser
	if err := c.ShouldBindJSON(&newUser); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := newUser.Create(app.DB)
	if err != nil {
		app.ch_error <- err
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	app.ch_user <- &newUser

	c.JSON(http.StatusCreated, gin.H{"inserted_id": newUser.Id})
}

func (app *App) GetUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		log.Println("Empty account id")
	}
	user_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var account models.User
	err = account.Get(app.DB, "", user_id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, account)
}

func (app *App) Login(c *gin.Context) {
	var loginPayload models.LoginPayload
	if err := c.ShouldBindJSON(&loginPayload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var account models.User
	err := account.Get(app.DB, loginPayload.Email, -1)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	hash, err := loginPayload.Login(account.Password, account.Id)
	if err != nil {
		app.ch_error <- err
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	app.ch_login <- account.Id
	c.JSON(http.StatusOK, gin.H{"access_token": hash})
}
