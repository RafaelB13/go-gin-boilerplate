package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rafaelb13/go-gin-boilerplate/config"
	"github.com/rafaelb13/go-gin-boilerplate/handler"
	"github.com/rafaelb13/go-gin-boilerplate/models"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

type Controller struct{}

func UserController() *Controller {
	return &Controller{}
}

// Index is the handler for the /user route
func (c *Controller) Index(ctx *gin.Context) {
	// Create a new user with sample information
	user := []models.User{}

	db = config.GetDatabase()
	logger = config.GetLogger("controller")

	if err := db.Find(&user).Error; err != nil {
		handler.SendError(ctx, http.StatusInternalServerError, "error listing users")
	}

	// Send a success response with user data
	handler.SendSuccess(ctx, "index-user", user)
}

func (c *Controller) Create(ctx *gin.Context) {
	request := models.User{}

	db = config.GetDatabase()
	logger = config.GetLogger("controller")

	ctx.BindJSON(&request)

	user := models.User{
		ID:       uuid.New(),
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	if err := db.Create(&user).Error; err != nil {
		logger.Errorf("ERROR CREATING OPENING: %v", err.Error())
		handler.SendError(ctx, http.StatusInternalServerError, "error user opening on database")
	}

	handler.SendSuccess(ctx, "create-user", user)

}
