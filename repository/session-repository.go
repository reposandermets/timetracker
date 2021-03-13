package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/reposandermets/timetracker/entity"
)

type SessionRepository interface {
	Save(session entity.Session) entity.Session
	Update(session entity.Session) entity.Session
	FindAll() []entity.Session
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewSessionRepository() SessionRepository {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&entity.Session{}, &entity.User{})

	result := db.First(&entity.User{})
	if result.RowsAffected == 0 {
		db.Create(&entity.User{
			Name: "Admin",
		})
	}
	return &database{
		connection: db,
	}
}

func (db *database) CloseDB() {
	db.connection.Close()
}

func (db *database) Save(session entity.Session) entity.Session {
	db.connection.Create(&session)
	return session
}

func (db *database) Update(session entity.Session) entity.Session {
	db.connection.Save(&session)
	return session
}

func (db *database) FindAll() []entity.Session {
	var sessions []entity.Session
	db.connection.Find(&sessions)
	return sessions
}
