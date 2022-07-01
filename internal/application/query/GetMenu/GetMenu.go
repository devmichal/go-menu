package GetMenu

import (
	"github.com/devmichal/foods/internal/application/query"
	"github.com/devmichal/foods/internal/domain"
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
	idMenu := c.Params("id")
	menu := h.repository.GetMenu(idMenu)

	queryDTO := query.MenuQuery{
		Pricing: menu.Price,
		Menu:    menu.Menu,
	}

	return c.JSON(queryDTO)
}
