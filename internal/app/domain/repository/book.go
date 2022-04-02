package repository

import "bookinfo/graph/model"

type IBook interface {
	GetAllBooks() ([]*model.Book, error)
}
