package usecase

import (
	"bookinfo/graph/model"
	"bookinfo/internal/app/domain/repository"
)

func GetAuthorsByNameAndAllRespectiveBooks(repo repository.IAuthor, name *string) ([]*model.Author, error) {
	res, err := repo.GetAuthorByName(name)
	return res, err
}
