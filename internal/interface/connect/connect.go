package connect

import (
	"github.com/devmichal/foods/cmd/flags"
	"github.com/devmichal/foods/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (db *gorm.DB) {
	db, err := gorm.Open(postgres.Open(flags.DB_CONFIG), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func Migrate() {
	db := Connect()

	db.AutoMigrate(&domain.Menu{})
}
