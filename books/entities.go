package books

import (
	"time"
)

type (
	Book struct {
		ID          int       `json:"id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Author      string    `json:"author"`
		Edition     int       `json:"edition"`
		ShelfID     int       `db:"shelf_id" json:"-"`
		CreatedAt   time.Time `json:"-"`
		UpdatedAt   time.Time `json:"-"`
		BookShelf   Shelf     `json:"shelf" ref:"shelf_id" fk:"id"`
	}

	Shelf struct {
		ID       int    `json:"id"`
		Capacity int    `json:"-"`
		Amount   int    `json:"-"`
		Books    []Book `json:"-" ref:"id" fk:"shelf_id"`
	}
)

func (Book) Table() string {
	return "book"
}

func (Shelf) Table() string {
	return "shelf"
}
