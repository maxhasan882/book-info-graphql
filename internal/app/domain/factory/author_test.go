package factory

import (
	"bookinfo/graph/model"
	"bookinfo/internal/app/adapter/db"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestAuthorGenerate(t *testing.T) {
	data := []*db.Author{
		{ID: 1, Name: "some", Biography: "hello", Books: []*db.Book{
			{ID: 1, Price: 12, Title: "hello"},
		}},
	}

	expected := []*model.Author{
		{ID: "1", Name: "some", Biography: "hello", Books: []*model.Book{
			{ID: "1", Price: 12, Title: "hello"},
		}},
	}
	response := Author{}.Generate(data)
	assert.Equal(t, expected, response)
}
