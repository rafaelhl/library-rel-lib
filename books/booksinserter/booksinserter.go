//go:generate mockery -name=InserterEngine
//go:generate mockery -name=ShelfFinder

package booksinserter

import (
	"context"
	"errors"

	"github.com/rafaelhl/library-rel-lib/books"
)

type (
	InserterEngine interface {
		InsertBook(ctx context.Context, book books.Book) error
	}

	ShelfFinder interface {
		FindShelf(ctx context.Context, shelfID int) (books.Shelf, error)
	}

	Inserter struct {
		inserterEngine InserterEngine
		shelfFinder    ShelfFinder
	}
)

func New(repository InserterEngine, shelfFinder ShelfFinder) Inserter {
	return Inserter{
		inserterEngine: repository,
		shelfFinder:    shelfFinder,
	}
}

func (i Inserter) InsertBook(ctx context.Context, book books.Book) error {
	shelf, err := i.shelfFinder.FindShelf(ctx, book.BookShelf.ID)
	if err != nil {
		return err
	}
	if len(shelf.Books) >= shelf.Capacity {
		return errors.New("shelf fully choose another shelf")
	}

	return i.inserterEngine.InsertBook(ctx, book)
}
