package factory

import (
	"bookinfo/graph/model"
	"bookinfo/internal/app/adapter/db"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestBookGenerate(t *testing.T) {
	data := []*db.Book{
		{ID: 1, Title: "some", Price: 12, IsbnNo: false, Authors: []*db.Author{{ID: 1, Name: "Hello", Biography: "test"}}},
	}

	expected := []*model.Book{
		{ID: "1", Title: "some", Price: 12, IsbnNo: false, Authors: []*model.Author{{ID: "1", Name: "Hello", Biography: "test"}}},
	}
	response := Book{}.Generate(data)
	assert.Equal(t, expected, response)
}
