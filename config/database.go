package config

import (
	"fmt"
	"os"

	"github.com/rafaelb13/go-gin-boilerplate/models"
	"github.com/rafaelb13/go-gin-boilerplate/utils"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func envInfos() utils.DBConnection {
	DotEnv().LoadEnv()

	// Retorne uma interface anônima contendo os dados
	return utils.DBConnection{
		DBSelect: os.Getenv("DB_SELECT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
	}
}

func initDB() (*gorm.DB, error) {

	var dns string

	if env := envInfos(); env.DBSelect == "mysql" {
		dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			env.Username,
			env.Password,
			env.Host,
			env.Port,
			env.Name,
		)
	} else if env.DBSelect == "postgres" {
		dns = fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
			env.Username,
			env.Name, // Certifique-se de ter a propriedade DBName configurada no ambiente
			env.Password,
			env.Host,
			env.Port,
		)
	}

	var db *gorm.DB
	var err error

	if env := envInfos(); env.DBSelect == "mysql" {
		db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	} else if env.DBSelect == "postgres" {
		db, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	} else if env.DBSelect == "mongodb" {
	}

	if err != nil {
		logger.Error("Error: ", err)
	}

	//Migrate the Schema
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		logger.Errorf("DB Automigration error: %v", err)
		return nil, err
	}

	logger.Info("Database Connection has successful")
	return db, nil

}
