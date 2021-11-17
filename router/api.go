package router

import (
	system "github.com/Hye-Ararat/Lava/system/information"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	v1 := api.Group("/v1")
	systemV1 := v1.Group("/system")
	systemV1.Get("/", system.GetSystemInformation)
}
