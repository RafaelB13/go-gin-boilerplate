package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pborman/uuid"
	"github.com/rafaelb13/go-gin-boilerplate/handler"
	"github.com/rafaelb13/go-gin-boilerplate/models"
)

type Controller struct{}

func UserController() *Controller {
	return &Controller{}
}

// Index is the handler for the /user route
func (c *Controller) Index(ctx *gin.Context) {
	// Create a new user with sample information
	user := models.User{
		Id:    uuid.New(),
		Name:  "Rafael Borges",
		Email: "rafaelborgesdev@gmail.com", // Contact me 🙂
	}

	// Send a success response with user data
	handler.SendSuccess(ctx, "index-user", user)
}
