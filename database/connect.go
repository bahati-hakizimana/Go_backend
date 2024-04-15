package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
		
	}
	// no error

	dsn:=os.Getenv("DSN")
	database,err:=gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil{
		fmt.Println("could not connect to database")
	}
	DB=database
}