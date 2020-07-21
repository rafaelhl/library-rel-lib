package repository

import (
	"context"

	"github.com/Fs02/rel"
	"github.com/Fs02/rel/where"

	"github.com/rafaelhl/library-rel-lib/books"
)

type BooksRepository struct {
	repo rel.Repository
}

func New(repo rel.Repository) BooksRepository {
	return BooksRepository{
		repo: repo,
	}
}

func (r BooksRepository) FindShelf(ctx context.Context, shelfID int) (books.Shelf, error) {
	var shelf books.Shelf
	err := r.repo.Find(ctx, &shelf, where.Eq("id", shelfID))
	if err != nil {
		return books.Shelf{}, err
	}
	err = r.repo.Preload(ctx, &shelf, "books")
	return shelf, err
}

func (r BooksRepository) InsertBook(ctx context.Context, book books.Book) error {
	return r.repo.Transaction(ctx, func(ctx context.Context) error {
		book.ShelfID = book.BookShelf.ID
		r.repo.MustInsert(ctx, &book, rel.Cascade(false))
		return r.repo.Update(ctx, &book.BookShelf, rel.Inc("amount"), rel.Reload(false))
	})
}
