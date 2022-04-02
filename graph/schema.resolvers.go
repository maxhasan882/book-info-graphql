package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bookinfo/graph/generated"
	"bookinfo/graph/model"
	"bookinfo/internal/app/adapter/db"
	dbConn "bookinfo/internal/app/adapter/db/connections"
	"context"
	"log"
	"strconv"
)

func (r *queryResolver) Author(ctx context.Context, name *string) (*model.Author, error) {
	database, err := dbConn.GetDatabase()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var author *db.Author
	database.Preload("Books").Where("name = ?", name).First(&author)
	var authorResponse = &model.Author{ID: strconv.Itoa(int(author.ID)), Name: author.Name, Biography: author.Biography}
	var books []*model.Book
	for _, book := range author.Books {
		books = append(books, &model.Book{ID: strconv.Itoa(int(book.ID)), Title: book.Title, Price: book.Price,
			IsbnNo: book.IsbnNo})
	}
	authorResponse.Books = books
	return authorResponse, nil
}

func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	database, err := dbConn.GetDatabase()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var books []*db.Book
	database.Preload("Authors").Find(&books)
	var booksResponse []*model.Book
	for _, book := range books {
		var authors []*model.Author
		for _, author := range book.Authors {
			authors = append(authors, &model.Author{ID: strconv.Itoa(int(author.ID)), Name: author.Name,
				Biography: author.Biography})
		}
		booksResponse = append(booksResponse, &model.Book{ID: strconv.Itoa(int(book.ID)), Title: book.Title,
			IsbnNo: book.IsbnNo, Price: book.Price, Authors: authors})
	}
	return booksResponse, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
