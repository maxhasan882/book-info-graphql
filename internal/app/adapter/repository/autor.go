package repository

import (
	"bookinfo/graph/model"
	"bookinfo/internal/app/adapter/db"
	dbConn "bookinfo/internal/app/adapter/db/connections"
	"bookinfo/internal/app/domain/factory"
	"log"
)

type Author struct{}

func (a Author) GetAuthorByName(name *string) (*model.Author, error) {
	database, err := dbConn.GetDatabase()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var author *db.Author
	database.Preload("Books").Where("name = ?", name).First(&author)
	authorFactory := factory.Author{}
	return authorFactory.Generate(author), nil
}
