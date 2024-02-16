package main

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var handler = func(c *fiber.Ctx) error { return nil }

func main() {
	app := fiber.New()

	app.Get("/john/:age", handler)
	app.Post("/register", handler)

	data, _ := json.MarshalIndent(app.Stack(), "", "  ")
	fmt.Println(string(data))

	app.Listen(":3000")
}
