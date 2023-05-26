package main

import (
	"fmt"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/pbnjay/memory"
)

func main() {
	app := fiber.New()

	memStats := &runtime.MemStats{}
	runtime.ReadMemStats(memStats)

	totalMemory := memStats.TotalAlloc / 1024
	fmt.Printf("Memoria Total: %d KB \n", totalMemory)

	fmt.Println("Total de memoria RAM: ", memory.TotalMemory())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("Hola, mundo")
	})

	app.Listen(":3000")
}
