//go:generate mockery -name=BookListFindEngine

package booklistfinder

import (
	"context"

	"github.com/rafaelhl/library-rel-lib/books"
)

type (
	BookListFindEngine interface {
		FindAllBooks(ctx context.Context) ([]books.Book, error)
	}

	Finder struct {
		engine BookListFindEngine
	}
)

func New(engine BookListFindEngine) Finder {
	return Finder{
		engine: engine,
	}
}

func (f Finder) GetAllBooks(ctx context.Context) ([]books.Book, error) {
	return f.engine.FindAllBooks(ctx)
}
