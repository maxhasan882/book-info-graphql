package repository

import (
	"bookinfo/graph/model"
	"bookinfo/internal/app/adapter/db"
	dbConn "bookinfo/internal/app/adapter/db/connections"
	"bookinfo/internal/app/domain/factory"
	"log"
)

type Book struct{}

// GetAllBooks Get all Book with all associate Author list
func (b Book) GetAllBooks() ([]*model.Book, error) {
	database, err := dbConn.GetDatabase()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var books []*db.Book
	database.Preload("Authors").Find(&books)
	bookFactory := factory.Book{}
	return bookFactory.Generate(books), nil
}
