package config

import (
	"fmt"
	"os"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error
	logger = GetLogger("config")
	// Initialize MySQL
	db, err = initMySQL()
	if err != nil {
		return fmt.Errorf("error initializing mysql: %v", err)
	}
	return nil
}

func ServerListenner() string {
	DotEnv().LoadEnv()

	appURL := os.Getenv("APP_URL")
	appURL = func() string {
		if appURL != "" {
			return appURL
		}
		return ":"
	}()

	port := os.Getenv("APP_PORT")
	port = func() string {
		if port != "" {
			return ":" + port // Adicione ":" antes da porta
		}
		return ":8888" // Porta padrão, se não estiver definida no arquivo .env
	}()

	return appURL + port

}

func GetLogger(p string) *Logger {
	// Initialize Logger
	logger = NewLogger(p)
	return logger
}

func GetDatabase() *gorm.DB {
	return db
}
