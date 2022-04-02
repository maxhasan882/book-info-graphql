package db

type Author struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Biography string  `json:"biography"`
	Books     []*Book `gorm:"many2many:authors_books" json:"books"`
}

type Book struct {
	ID      uint      `json:"id"`
	Title   string    `json:"title"`
	Price   float64   `json:"price"`
	IsbnNo  bool      `json:"isbn_no"`
	Authors []*Author `gorm:"many2many:authors_books" json:"books"`
}
