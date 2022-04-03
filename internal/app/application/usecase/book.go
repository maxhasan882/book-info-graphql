package usecase

import (
	"bookinfo/graph/model"
	"bookinfo/internal/app/domain/repository"
)

func GetAllBooksWithRespectiveAuthors(repo repository.IBook) ([]*model.Book, error) {
	res, err := repo.GetAllBooks()
	return res, err
}
