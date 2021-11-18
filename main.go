package main

import (
	"github.com/Hye-Ararat/Lava/router"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()
	router.RegisterRoutes(app)
	log.Fatal(app.Listen("0.0.0.0:2221"))
}
