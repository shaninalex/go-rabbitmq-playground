package app

import (
	"account/app/models"
	"log"
	"net/http"

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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println(newUser)

	c.JSON(http.StatusCreated, gin.H{"success": "ok"})
}

//
// func (app *App) GetUser(c *gin.Context) {
// 	id := c.Param("id")
// 	if id == "" {
// 		log.Println("Empty account id")
// 	}
//
// 	var account models.GetCreateAccount
// 	account.Sub = id
// 	err := account.Get(app.Collection)
// 	if err != nil {
// 		log.Println(err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"message": "User does not exists"})
// 		return
// 	}
//
// 	c.JSON(http.StatusOK, account)
// }
//
// func (app *App) UpdateUser(c *gin.Context) {
// 	id := c.Param("id")
// 	if id == "" {
// 		log.Println("Empty account id")
// 	}
//
// 	var payload models.UpdateAccount
// 	if err := c.ShouldBindJSON(&payload); err != nil {
// 		log.Println(err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
// 		return
// 	}
//
// 	if reflect.ValueOf(payload).IsZero() {
// 		log.Println(fmt.Errorf("payload %v is empty or contain wrong values. Nothing to udpate", payload))
// 		c.JSON(http.StatusInternalServerError, gin.H{"message": "Cant handle payload (empty or incorrect)"})
// 		return
// 	}
//
// 	err := models.Update(app.Collection, id, payload)
// 	if err != nil {
// 		log.Println(err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to update account"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"success": true})
// }
