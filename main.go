package main

import (
	"log"
	"os"

	"github.com/ernestokorpys/gobackend/database"
	"github.com/ernestokorpys/gobackend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	database.Connect()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error")
	}
	port := os.Getenv("PORT")
	app := fiber.New() // creat una instacncia en la aplicacion go
	//esta bien el orden ya que es como la carga de como se van a manejar los datos y el otro el el que se pone a escuchar
	routes.Setup(app)      //maneja las solicitudes entrantes
	app.Listen(":" + port) //Escucha las solicitudes del puerto constantemente
}
