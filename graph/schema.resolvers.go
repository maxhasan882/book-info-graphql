package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bookinfo/graph/generated"
	"bookinfo/graph/model"
	"bookinfo/internal/app/adapter/repository"
	"bookinfo/internal/app/application/usecase"
	"context"
)

func (r *queryResolver) Author(ctx context.Context, name *string) (*model.Author, error) {
	return usecase.GetAuthorByIdAndAllRespectiveBooks(repository.Author{}, name)
}

func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	return usecase.GetAllBooksWithRespectiveAuthors(repository.Book{})
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
