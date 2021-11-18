package router

import (
	serverRouter "github.com/Hye-Ararat/Lava/router/server"
	system "github.com/Hye-Ararat/Lava/system/information"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	systemAPI := api.Group("/system")
	systemAPI.Get("/", system.GetSystemInformation)

	serverAPI := api.Group("/servers/:server")
	serverAPI.Get("/state", serverRouter.GetState)
}
