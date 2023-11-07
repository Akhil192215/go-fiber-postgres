package config

import (
	"log"
	"os"

	"github.com/Akhil192215/go-fiber/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	var err error
	DBURL := os.Getenv("DBURL")
	dsn := DBURL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect to database! \n", err.Error())
		os.Exit(2)
	}
	log.Println("Connected to database successfully ")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	Database = DbInstance{Db: db}
}
