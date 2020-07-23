package bookfinder_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rafaelhl/library-rel-lib/books"
	"github.com/rafaelhl/library-rel-lib/books/bookfinder"
	"github.com/rafaelhl/library-rel-lib/books/bookfinder/mocks"
)

var (
	bookID   = 1
	expected = books.Book{
		ID:          bookID,
		Title:       "Livro de Teste",
		Description: "Esse livro é de teste",
		Author:      "Rafael Holanda",
		Edition:     1,
		BookShelf: books.Shelf{
			ID: 1,
		},
	}
)

func TestFinder_FindBook(t *testing.T) {
	ctx := context.Background()
	fetcher := new(mocks.BookFetcher)
	fetcher.On("FindBookByID", ctx, bookID).Return(expected, nil)
	finder := bookfinder.New(fetcher)

	book, err := finder.FindBook(ctx, bookID)

	assert.NoError(t, err)
	assert.Equal(t, expected, book)
	fetcher.AssertExpectations(t)
}

func TestFinder_FindBookFail(t *testing.T) {
	ctx := context.Background()
	fetcher := new(mocks.BookFetcher)
	fetcher.On("FindBookByID", ctx, bookID).Return(books.Book{}, errors.New("unexpected error|"))
	finder := bookfinder.New(fetcher)

	_, err := finder.FindBook(ctx, bookID)

	assert.Error(t, err)
}
