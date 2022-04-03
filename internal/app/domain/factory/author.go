package factory

import (
	"bookinfo/graph/model"
	"bookinfo/internal/app/adapter/db"
	"strconv"
)

type Author struct {
}

// Generate Serialize Author data
func (t Author) Generate(authors []*db.Author) []*model.Author {
	var authorsData []*model.Author
	for _, author := range authors {
		var books []*model.Book
		for _, book := range author.Books {
			books = append(books, &model.Book{ID: strconv.Itoa(int(book.ID)), Title: book.Title, Price: book.Price,
				IsbnNo: book.IsbnNo})
		}
		authorsData = append(authorsData, &model.Author{ID: strconv.Itoa(int(author.ID)),
			Name: author.Name, Biography: author.Biography, Books: books})
	}
	return authorsData
}
