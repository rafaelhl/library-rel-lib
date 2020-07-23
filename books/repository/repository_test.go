package repository_test

import (
	"context"
	"math/rand"
	"testing"

	"github.com/Fs02/rel"
	"github.com/Fs02/rel/reltest"
	"github.com/Fs02/rel/where"
	"github.com/stretchr/testify/assert"

	"github.com/rafaelhl/library-rel-lib/books"
	"github.com/rafaelhl/library-rel-lib/books/repository"
)

var (
	expectedShelf = books.Shelf{
		ID:       1,
		Capacity: 2,
		Amount:   0,
	}
	expectedBooks = []books.Book{
		{
			ID:          1,
			Title:       "Teste 1",
			Description: "Teste 1",
			Author:      "Teste 1",
			Edition:     rand.Int(),
			ShelfID:     1,
		},
		{
			ID:          2,
			Title:       "Teste 2",
			Description: "Teste 2",
			Author:      "Teste 2",
			Edition:     rand.Int(),
			ShelfID:     1,
		},
	}
)

func TestBooksRepository_FindShelf(t *testing.T) {
	repo := reltest.New()
	booksRepository := repository.New(repo)
	shelfID := 1

	repo.ExpectFind(where.Eq("id", shelfID)).Result(expectedShelf)
	repo.ExpectPreload("books").ForType("books.Shelf").Result(expectedBooks)

	shelf, err := booksRepository.FindShelf(context.Background(), shelfID)

	repo.AssertExpectations(t)

	assert.NoError(t, err)
	assert.Equal(t, books.Shelf{
		ID:       expectedShelf.ID,
		Capacity: expectedShelf.Capacity,
		Amount:   expectedShelf.Amount,
		Books:    expectedBooks,
	}, shelf)
}

func TestBooksRepository_InsertBook(t *testing.T) {
	repo := reltest.New()
	booksRepository := repository.New(repo)
	book := books.Book{
		Title:       "Livro de Teste",
		Description: "Esse livro é de teste",
		Author:      "Rafael Holanda",
		Edition:     1,
		BookShelf: books.Shelf{
			ID: 1,
		},
	}

	repo.ExpectTransaction(func(repo *reltest.Repository) {
		repo.ExpectInsert(rel.Cascade(false)).ForType("books.Book")
		repo.ExpectUpdate(rel.Inc("amount"), rel.Reload(false)).ForType("books.Shelf")
	})

	err := booksRepository.InsertBook(context.Background(), book)

	repo.AssertExpectations(t)
	assert.NoError(t, err)
}

func TestBooksRepository_FindBookByID(t *testing.T) {
	repo := reltest.New()
	booksRepository := repository.New(repo)
	expected := expectedBooks[0]

	repo.ExpectFind(where.Eq("id", expected.ID)).Result(expected)
	repo.ExpectPreload("book_shelf").For(&expected).Result(expectedShelf)

	book, err := booksRepository.FindBookByID(context.Background(), expected.ID)

	assert.NoError(t, err)
	assert.Equal(t, books.Book{
		ID:          expected.ID,
		Title:       expected.Title,
		Description: expected.Description,
		Author:      expected.Author,
		Edition:     expected.Edition,
		ShelfID:     expected.ShelfID,
		BookShelf:   expectedShelf,
	}, book)
}
