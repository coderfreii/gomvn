package database

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/gomvn/gomvn/internal/config"
	"github.com/gomvn/gomvn/internal/entity"
)

func New() (*gorm.DB, error) {
	if err := os.MkdirAll(config.Dirname, os.ModeDir); err != nil {
		log.Println("cannot create data directory")
		return nil, err
	}

	db, err := gorm.Open("sqlite3", config.Dirname+"/data.db")
	if err != nil {
		return nil, err
	}

	// db.LogMode(true)
	db.AutoMigrate(&entity.User{}, &entity.Path{})
	db.Model(&entity.Path{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	return db, nil
}
