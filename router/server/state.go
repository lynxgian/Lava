package serverRouter

import (
	"github.com/Hye-Ararat/Lava/server"
	"github.com/gofiber/fiber/v2"
)

func GetState(c *fiber.Ctx) error {
	status, err := server.GetState(c.Params("server"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}
	errSend := c.JSON(fiber.Map{
		"status": "success",
		"data":   status,
	})
	if errSend != nil {
		return errSend
	}
	return nil
}
