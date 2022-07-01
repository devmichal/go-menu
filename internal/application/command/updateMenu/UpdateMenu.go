package updateMenu

import (
	"github.com/devmichal/foods/internal/application/command"
	"github.com/devmichal/foods/internal/domain"
	"github.com/devmichal/foods/internal/interface/db"
	"github.com/devmichal/foods/internal/validator"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	repository domain.MenuRepository
}

func UpdateMenu(repository domain.MenuRepository) *Handler {
	return &Handler{
		repository: repository,
	}
}

func (h *Handler) Menu(c *fiber.Ctx) error {
	dto := new(command.Menu)
	idMenu := c.Params("id")

	if err := c.BodyParser(dto); err != nil {
		return err
	}

	errors := validator.ValidateStruct(*dto)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	db.CreateMenu(*dto)

	return c.JSON(idMenu)
}
