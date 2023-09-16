package main

import (
	"github.com/rafaelb13/go-gin-boilerplate/config"
	app "github.com/rafaelb13/go-gin-boilerplate/service"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//Init Config
	err := config.Init()
	if err != nil {
		logger.Errorf("error starting settings %s", err.Error())
		return
	}
	// Start
	app.ServerInitialize()
}
