package main

import (
	"github.com/eneasdev5/go-fiber-app/src/database"
	"github.com/eneasdev5/go-fiber-app/src/domain"
	"github.com/eneasdev5/go-fiber-app/src/repository/mysql"
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
	dbConnect := database.Connect()
	repository := mysql.NewMysqlDBRepositoryBook(dbConnect)

	// define the engine views
	engine := django.New("./src/views", ".html")

	// new instance struct fiber
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// routes to app
	app.Get("/books", func(c *fiber.Ctx) error {
		books, err := repository.GetAll()
		if err != nil {
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"books": books,
		})
	})

	app.Post("/books", func(c *fiber.Ctx) error {
		b := domain.Book{
			Title:       "Test Title",
			Body:        "Hello Worl 1000",
			Description: "There are many variations of passages of Lo",
		}

		book, err := repository.Store(b)
		if err != nil {
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"book": book,
		})
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"title": "Page Index",
		})
	})

	app.Get("/params", func(c *fiber.Ctx) error {
		d := new(Dados)

		if err := c.QueryParser(d); err != nil {
			return err
		}

		return c.Render("home", fiber.Map{
			"title":     "page Home",
			"queryID":   c.Query("id"),
			"queryNome": c.Query("nome"),
			"parse":     d,
			"queries":   c.Queries(),
		})
	})

	// expose resource files static
	app.Static("static", "./public")

	// log.Info("[message: listen server in port 3000] ->> [http://localhost:3000] or [http://127.0.0.1:3000]")
	// listen server in port 3000 ->> http://localhost:3000 or http://127.0.0.1:3000
	log.Fatal(app.Listen(":3000"))
}
