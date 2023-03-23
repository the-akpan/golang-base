package models

import (
	"os"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// database, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
	var dbConn string
	if dbConn = os.Getenv("SQLITE"); dbConn == "" {
		dbConn = "test.db"
	}

	database, err := gorm.Open(sqlite.Open(dbConn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	database.AutoMigrate(&User{})
	DB = database
}

func ConnectDatabaseMock() {
	var dbConn string
	if dbConn = os.Getenv("SQLITE"); dbConn == "" {
		dbConn = "test.db"
	}

	database, err := gorm.Open(sqlite.Open(dbConn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	database.AutoMigrate(&User{})
	DB = database
}

type Base struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time `json:"createdat" gorm:"index:"`
	UpdatedAt time.Time `json:"updatedat" gorm:"index:"`
	CreatedBy uuid.UUID `json:"createdby" gorm:"type:uuid;index:"`
	UpdatedBy uuid.UUID `json:"updatedby" gorm:"type:uuid;index:"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) error {
	if base.ID.String() == "00000000-0000-0000-0000-000000000000" {
		uuid := uuid.NewV4()
		tx.Statement.SetColumn("ID", uuid)
	}
	tx.Statement.SetColumn("CreatedAt", time.Now())
	tx.Statement.SetColumn("UpdatedAt", time.Now())
	return nil
}

func (base *Base) BeforeUpdate(tx *gorm.DB) error {
	tx.Statement.SetColumn("UpdatedAt", time.Now())
	return nil
}
