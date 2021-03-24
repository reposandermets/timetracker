package repository

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/reposandermets/timetracker/entity"
)

type SessionRepository interface {
	FindSessionById(session *entity.Session, SessionID uuid.UUID) *gorm.DB
	FindUserById(user *entity.User, UserID uuid.UUID) *gorm.DB
	FindStartedSessionByUserId(session *entity.Session, UserID uuid.UUID) *gorm.DB
	Save(session *entity.Session) *gorm.DB
	Update(session *entity.Session) *gorm.DB
	FindSessionsByUserId(sessions *[]entity.Session, UserID uuid.UUID) *gorm.DB
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

func (db *database) FindSessionsByUserId(sessions *[]entity.Session, UserID uuid.UUID) *gorm.DB {
	return db.connection.Where("user_id = ?", UserID.String()).Order("started_at desc").Find(sessions)
}

func (db *database) CloseDB() {
	db.connection.Close()
}

func (db *database) FindSessionById(session *entity.Session, SessionID uuid.UUID) *gorm.DB {
	return db.connection.First(session, "id = ?", SessionID.String())
}

func (db *database) FindUserById(user *entity.User, UserID uuid.UUID) *gorm.DB {
	return db.connection.First(user, "id = ?", UserID.String())
}

func (db *database) Save(session *entity.Session) *gorm.DB {
	return db.connection.Create(session)
}

func (db *database) Update(session *entity.Session) *gorm.DB {
	return db.connection.Save(session)
}

func (db *database) FindStartedSessionByUserId(session *entity.Session, UserID uuid.UUID) *gorm.DB {
	return db.connection.Where("user_id = ? AND status = ?", UserID.String(), "started").First(session)
}
