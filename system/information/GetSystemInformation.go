package system

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pbnjay/memory"
	"runtime"
)

type Information struct {
	Version string
	OS      string
	CPU     CPU
	Memory  Memory
}
type CPU struct {
	Cores int
}
type Memory struct {
	Total uint64
}

func GetSystemInformation(c *fiber.Ctx) error {
	info := &Information{
		Version: "0.0.1",
		OS:      runtime.GOOS + " " + runtime.GOARCH,
		CPU: CPU{
			Cores: runtime.NumCPU(),
		},
		Memory: Memory{
			Total: memory.TotalMemory(),
		},
	}
	return c.JSON(&fiber.Map{
		"status": "success",
		"data":   info,
	})
}
