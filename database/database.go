package database

import (
	"log"
	"pokemonkanto/models"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func ConnectDB() {

	//connect to db
	var err error
	db, err = gorm.Open("sqlite3", "./pokemon.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	//creates structs in db tables
	db.AutoMigrate(&models.FeedbackEntry{})

}
