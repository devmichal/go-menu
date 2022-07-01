package db

import (
	"github.com/devmichal/foods/internal/application/command"
	"github.com/devmichal/foods/internal/domain"
	"github.com/devmichal/foods/internal/interface/connect"
	"gorm.io/gorm"
)

func CloseConnection(orm *gorm.DB) {

	db := connect.Connect()

	// Create
	db.Create(&domain.Menu{Menu: "chiec", Price: 100})
}

func CreateMenu(dto command.Menu) {
	db := connect.Connect()
	db.Create(&domain.Menu{Menu: dto.Menu, Price: dto.Price})
}
