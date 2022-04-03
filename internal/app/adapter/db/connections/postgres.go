package db

import (
	"bookinfo/cmd/app/config"
	"bookinfo/internal/app/adapter/db"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type dbConfig struct {
	host     string
	port     string
	user     string
	dbname   string
	password string
}

// getDatabaseUrl format db url using environment variable and prepare postgres connection string
func getDatabaseUrl() string {
	var _config = dbConfig{config.DBHost, config.DBPort, config.DBUser, config.DBName, config.DBPassword}
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s",
		_config.host, _config.port, _config.user, _config.dbname, _config.password)
}

// GetDatabase Get database connection
func GetDatabase() (*gorm.DB, error) {
	_db, err := gorm.Open(postgres.Open(getDatabaseUrl()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	return _db, err
}

// RunMigrations Migrate database automatically
func RunMigrations(_db *gorm.DB) {
	err := _db.AutoMigrate(&db.Author{}, &db.Book{})
	if err != nil {
		log.Println(err)
	}
}
