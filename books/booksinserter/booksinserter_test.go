package booksinserter_test

import (
	"context"
	"errors"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rafaelhl/library-rel-lib/books"
	"github.com/rafaelhl/library-rel-lib/books/booksinserter"
	"github.com/rafaelhl/library-rel-lib/books/booksinserter/mocks"
)

var (
	book = books.Book{
		Title:       "Livro de Teste",
		Description: "Esse livro é de teste",
		Author:      "Rafael Holanda",
		Edition:     1,
		BookShelf: books.Shelf{
			ID: 1,
		},
	}
	shelf = books.Shelf{
		ID:       1,
		Capacity: 2,
		Amount:   0,
	}
)

type mocked struct {
	inserter *mocks.InserterEngine
	finder   *mocks.ShelfFinder
}

func (m mocked) assertExpectations(t *testing.T) {
	m.inserter.AssertExpectations(t)
	m.finder.AssertExpectations(t)
}

func doMock() (booksinserter.Inserter, mocked) {
	m := mocked{
		inserter: new(mocks.InserterEngine),
		finder:   new(mocks.ShelfFinder),
	}
	return booksinserter.New(m.inserter, m.finder), m
}

func TestInserter_InsertBook(t *testing.T) {
	ctx := context.Background()
	inserter, m := doMock()
	m.finder.On("FindShelf", ctx, book.BookShelf.ID).Return(shelf, nil)
	m.inserter.On("InsertBook", ctx, book).Return(nil)

	err := inserter.InsertBook(ctx, book)

	m.assertExpectations(t)
	assert.NoError(t, err)
}

func TestInserter_InsertBook_Fail(t *testing.T) {
	ctx := context.Background()
	inserter, m := doMock()
	m.finder.On("FindShelf", ctx, book.BookShelf.ID).Return(shelf, nil)
	m.inserter.On("InsertBook", ctx, book).Return(errors.New("unexpected error!"))

	err := inserter.InsertBook(ctx, book)

	m.assertExpectations(t)
	assert.Error(t, err)
}

func TestInserter_InsertBook_FindShelfFail(t *testing.T) {
	ctx := context.Background()
	inserter, m := doMock()
	m.finder.On("FindShelf", ctx, book.BookShelf.ID).Return(books.Shelf{}, errors.New("unexpected error!"))

	err := inserter.InsertBook(ctx, book)

	m.assertExpectations(t)
	assert.Error(t, err)
}

func TestInserter_InsertBook_ShelfFully(t *testing.T) {
	ctx := context.Background()
	inserter, m := doMock()
	s := books.Shelf{
		ID:       shelf.ID,
		Capacity: shelf.Capacity,
		Amount:   shelf.Amount,
		Books: []books.Book{
			{
				ID:          1,
				Title:       "Teste 1",
				Description: "Teste 1",
				Author:      "Teste 1",
				Edition:     rand.Int(),
				ShelfID:     shelf.ID,
			},
			{
				ID:          2,
				Title:       "Teste 2",
				Description: "Teste 2",
				Author:      "Teste 2",
				Edition:     rand.Int(),
				ShelfID:     shelf.ID,
			},
		},
	}
	m.finder.On("FindShelf", ctx, book.BookShelf.ID).Return(s, nil)

	err := inserter.InsertBook(ctx, book)

	m.assertExpectations(t)
	assert.Error(t, err)
}
