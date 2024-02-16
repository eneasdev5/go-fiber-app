package main

import (
	"strconv"

	"github.com/eneasdev5/go-fiber-app/src/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/template/django/v3"
)

type Dados struct {
	ID        int    `query:"id"`
	Nome      string `query:"nome"`
	Sobrenome string `query:"sobrenome"`
}

func main() {
	// define the engine views
	engine := django.New("../src/views", ".html")

	// new instance struct fiber
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// request database
	bookService := database.NewBook()

	// routes to app
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"title": "Page Index",
		})
	})

	app.Get("/books", func(c *fiber.Ctx) error {
		return c.Render("books", fiber.Map{
			"title": "Page Books",
			"books": bookService.GetAllBook(),
		})
	})

	app.Get("/home", func(c *fiber.Ctx) error {
		d := new(Dados)

		if err := c.QueryParser(d); err != nil {
			return err
		}

		log.Info("info")
		return c.Render("home", fiber.Map{
			"title":     "page Home",
			"queryID":   c.Query("id"),
			"queryNome": c.Query("nome"),
			"parse":     d,
			"queries":   c.Queries(),
		})
	})

	app.Get("/users/:id", func(c *fiber.Ctx) error {
		person := struct {
			ID   int    `json:"id"`
			Nome string `json:"nome"`
		}{5, "Eneas"}

		id := c.Params("id")
		if id != "" {
			paramId, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				return err
			}
			if paramId == int64(person.ID) {
				return c.JSON(person)
			}
		}

		c.Response().SetStatusCode(fiber.StatusNoContent)
		return c.JSON(nil)
	})

	// expose resource files static
	app.Static("static", "./public")

	log.Info("[message: listen server in port 3000] ->> [http://localhost:3000] or [http://127.0.0.1:3000]")
	// listen server in port 3000 ->> http://localhost:3000 or http://127.0.0.1:3000
	log.Fatal(app.Listen(":3000"))
}
