package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dbConfig struct {
	host     string
	port     int
	user     string
	dbname   string
	password string
}

var config = dbConfig{"localhost", 6753, "postgres", "test", "postgres"}

func getDatabaseUrl() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s",
		config.host, config.port, config.user, config.dbname, config.password)
}

func GetDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(getDatabaseUrl()), &gorm.Config{})
	return db, err
}

//func RunMigrations(db *gorm.DB) {
//	db.AutoMigrate(&Author{}, &Book{})
//}
