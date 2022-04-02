package db

import (
	"bookinfo/internal/app/adapter/db"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type dbConfig struct {
	host     string
	port     int
	user     string
	dbname   string
	password string
}

var config = dbConfig{"localhost", 5432, "postgres", "test", "mithu1996"}

func getDatabaseUrl() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s",
		config.host, config.port, config.user, config.dbname, config.password)
}

func GetDatabase() (*gorm.DB, error) {
	_db, err := gorm.Open(postgres.Open(getDatabaseUrl()), &gorm.Config{})
	return _db, err
}

func RunMigrations(_db *gorm.DB) {
	err := _db.AutoMigrate(&db.Author{}, &db.Book{})
	if err != nil {
		log.Println(err)
	}
}
