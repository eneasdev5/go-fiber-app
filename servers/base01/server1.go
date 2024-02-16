package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Software Developer with fiber")
	})

	app.Get("/config", func(c *fiber.Ctx) error {
		return c.JSON(app.Config())
	})

	// Parâmentros   www.meu-site.com/hello%20World
	app.Get("/index1/:value", func(c *fiber.Ctx) error {
		return c.SendString("Value: " + c.Params("value"))
	})

	// Parâmentros
	app.Get("/index2/:nome?", func(c *fiber.Ctx) error {
		if c.Params("nome") != "" {
			return c.SendString("Hello, " + c.Params("nome"))
		}
		return c.SendString("where is Jhon?")
	})

	// Parâmentros /*
	app.Get("/api/*", func(c *fiber.Ctx) error {
		return c.SendString("Api path: " + c.Params("*"))
	})

	// config files static
	app.Static("/", "./public")

	app.Listen(":3000")

}
