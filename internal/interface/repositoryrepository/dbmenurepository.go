package repositoryrepository

import (
	"github.com/devmichal/foods/internal/domain"
	"github.com/devmichal/foods/internal/interface/connect"
)

type InMemoryMenuRepository struct {
	food []domain.Menu
}

func NewInMemoryFoodRepository() *InMemoryMenuRepository {
	r := new(InMemoryMenuRepository)
	return r
}

func (r *InMemoryMenuRepository) SaveMenu(menu *domain.Menu) {
	db := connect.Connect()
	db.Create(menu)
}

func (r *InMemoryMenuRepository) UpdateMenu(menu *domain.Menu) {
	db := connect.Connect()
	db.Model(&menu).Update("id", 2)
}

func (r *InMemoryMenuRepository) GetMenu(idMenu string) domain.Menu {
	var menu domain.Menu
	db := connect.Connect()
	_ = db.First(&menu, idMenu)

	return menu
}
