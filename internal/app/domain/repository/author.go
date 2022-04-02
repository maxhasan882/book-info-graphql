package repository

import (
	"bookinfo/graph/model"
)

type IAuthor interface {
	GetAuthorByName(name *string) (*model.Author, error)
}
