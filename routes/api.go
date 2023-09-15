package routes

import (
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/rafaelb13/go-gin-boilerplate/controllers"
)

func ApiRoutes(r *gin.Engine) {
	// Initial route that returns the Go version in use
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"go_version": runtime.Version(),
		})
	})

	// Initialize the user controller
	user := controllers.UserController()
	// Route to access the Index handler in the user controller
	r.GET("/user", user.Index)
}
