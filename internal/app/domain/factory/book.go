package factory

import (
	"bookinfo/graph/model"
	"bookinfo/internal/app/adapter/db"
	"strconv"
)

type Book struct {
}

func (b Book) Generate(books []*db.Book) []*model.Book {
	var booksData []*model.Book
	for _, book := range books {
		var authors []*model.Author
		for _, author := range book.Authors {
			authors = append(authors, &model.Author{ID: strconv.Itoa(int(author.ID)), Name: author.Name,
				Biography: author.Biography})
		}
		booksData = append(booksData, &model.Book{ID: strconv.Itoa(int(book.ID)), Title: book.Title,
			IsbnNo: book.IsbnNo, Price: book.Price, Authors: authors})
	}
	return booksData
}
