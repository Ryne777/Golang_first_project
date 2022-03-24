package bdClient

import (
	"github.com/Ryne777/Golang_first_project/internal/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetConnectionDb(_ *config.Config) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("../../test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
