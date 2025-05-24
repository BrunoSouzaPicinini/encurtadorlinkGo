package repository

import (
	"encurtadorLink/src/model"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbConnection *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open("encurtador.db"))

	if err != nil {
		//Panic derruba aplicação
		// panic("failed to connect database")
		log.Fatal("failed to connect database")
		return
	}

	result := db.Select("select 1")

	if result.Error != nil {
		log.Fatal("failed to select")
		return
	}

	db.AutoMigrate(
		&model.Shortener{},
	)

	dbConnection = db
}
