package database

import (
	"log"
	"os"

	"github.com/ernestokorpys/gobackend/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load .env file")
	}
	dsn := os.Getenv("DSN")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database")
	} else {
		log.Println("Connected to the database successfully!")
	}
	DB = database
	// pointer := &DB
	// fmt.Printf("Pointer value: %p\n", pointer)
	// fmt.Printf("DB value: %v\n", *pointer)
	database.AutoMigrate( //create database topics
		&models.User{},
		&models.Blog{},
	)
}
