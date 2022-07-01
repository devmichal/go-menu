package cmd

import (
	"github.com/devmichal/foods/internal/application/command/createMenu"
	"github.com/devmichal/foods/internal/application/command/updateMenu"
	"github.com/devmichal/foods/internal/application/query/GetMenu"
	"github.com/devmichal/foods/internal/interface/repositoryrepository"
	"github.com/gofiber/fiber/v2"
)

func Route() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("Hello, World 👋! ")
	})

	repository := repositoryrepository.NewInMemoryFoodRepository()

	app.Post("/menu", createMenu.NewMenu(repository).Menu)
	app.Get("/menu/:id", GetMenu.NewMenu(repository).Menu)
	app.Put("/menu/:id", updateMenu.UpdateMenu(repository).Menu)

	return app
}
