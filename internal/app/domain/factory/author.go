package factory

import (
	"bookinfo/graph/model"
	"bookinfo/internal/app/adapter/db"
	"strconv"
)

type Author struct {
}

func (t Author) Generate(author *db.Author) *model.Author {
	var authorData = &model.Author{ID: strconv.Itoa(int(author.ID)), Name: author.Name, Biography: author.Biography}
	var books []*model.Book
	for _, book := range author.Books {
		books = append(books, &model.Book{ID: strconv.Itoa(int(book.ID)), Title: book.Title, Price: book.Price,
			IsbnNo: book.IsbnNo})
	}
	authorData.Books = books
	return authorData
}
