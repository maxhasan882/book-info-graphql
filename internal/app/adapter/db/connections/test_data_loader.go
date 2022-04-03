package db

import "bookinfo/internal/app/adapter/db"

func DataLoader() {
	authorBooks := []*db.Book{{Title: "Golang intro", Price: 120, IsbnNo: true}, {Title: "Learn graphql", Price: 299, IsbnNo: false}}
	authors := []*db.Author{{Name: "Hasan", Biography: "demo", Books: authorBooks}, {Name: "Milon", Biography: "test data", Books: authorBooks}}
	books := []*db.Book{{Title: "Golang intro", Price: 120, IsbnNo: true, Authors: authors}, {Title: "Learn graphql", Price: 299, IsbnNo: false, Authors: authors}}
	database, _ := GetDatabase()
	database.Create(&authors)
	database.Create(&books)
}
