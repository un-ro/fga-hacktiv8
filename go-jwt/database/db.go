package database

import (
	"fmt"
	"go-jwt/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const (
	HOST   = "localhost"
	PORT   = 5432
	USER   = "postgres"
	DBNAME = "fga"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	config := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", HOST, PORT, USER, DBNAME)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	fmt.Println("Database connection successfully opened")
	db.Debug().AutoMigrate(models.User{}, models.Product{})
}

func GetDB() *gorm.DB {
	return db
}
