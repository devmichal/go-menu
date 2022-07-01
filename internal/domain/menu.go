package domain

import (
	"github.com/devmichal/foods/internal/application/command"
	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	Menu  string  `json:"menu"`
	Price float32 `json:"price" gorm:'varchar(50)'`
}

func NewMenu(menu string, price float32) *Menu {
	return &Menu{
		Menu:  menu,
		Price: price,
	}
}

func FromDto(dto command.Menu) *Menu {
	return NewMenu(dto.Menu, dto.Price)
}
