package main

import (
	"log"
	"os"

	"github.com/bahati-hakizimana/Go_backend/database"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {

	database.Connect()

	err:=godotenv.Load()

	if  err != nil {

		log.Fatal("erro loading env")
		
	}

	port:=os.Getenv("PORT")
	app:=fiber.New()
	app.Listen(":" +port)

}