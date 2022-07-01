package createMenu

import (
	"github.com/devmichal/foods/internal/application/command"
	"github.com/devmichal/foods/internal/domain"
	"github.com/devmichal/foods/internal/validator"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	repository domain.MenuRepository
}

func NewMenu(repository domain.MenuRepository) *Handler {
	return &Handler{
		repository: repository,
	}
}

func (h *Handler) Menu(c *fiber.Ctx) error {
	dto := new(command.Menu)

	if err := c.BodyParser(dto); err != nil {
		return err
	}

	errors := validator.ValidateStruct(*dto)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	h.repository.SaveMenu(domain.FromDto(*dto))

	return c.JSON(dto.Menu)
}
