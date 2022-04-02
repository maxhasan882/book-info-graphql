package usecase

import (
	"bookinfo/internal/app/adapter/repository"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestAuthor(t *testing.T) {
	_, er := GetAllBooksWithRespectiveAuthors(repository.Book{})
	assert.Equal(t, er, nil)
}
