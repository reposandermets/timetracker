package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/reposandermets/timetracker/entity"
	"github.com/reposandermets/timetracker/server"
)

// import "github.com/reposandermets/timetracker/server"

func dbTest() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect database")
	}

	existingSession := entity.Session{}
	result := db.First(&existingSession, "id = ?", "8b8911f9-e09e-4c0e-bce0-b8b72f511053")
	println(result)
}

func main() {
	println("Hello world")
	dbTest()
	server.Run()
}
