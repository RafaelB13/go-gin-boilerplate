package config

import (
	"github.com/rafaelb13/go-gin-boilerplate/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initMySQL() (*gorm.DB, error) {

	dsn := "root:rafael@tcp(127.0.0.1:3306)/go-gin-dev?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("Error: ", err)
	}

	//Migrate the Schema
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		logger.Errorf("SQLite Automigration error: %v", err)
		return nil, err
	}

	logger.Info("Database Connection has successful")
	return db, nil

}
