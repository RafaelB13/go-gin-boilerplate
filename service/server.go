package service

import (
	"github.com/gin-gonic/gin"
	"github.com/rafaelb13/go-gin-boilerplate/config"
	"github.com/rafaelb13/go-gin-boilerplate/routes"
)

// ServerInitialize initializes and starts an HTTP server using the Gin framework.
// It creates a default Gin instance and runs it on port 8888.
func ServerInitialize() {

	// Create a default instance of the Gin server.
	app := gin.Default()

	routes.ApiRoutes(app)

	// Start the server on port 8888.
	app.Run(config.ServerListenner())
}
