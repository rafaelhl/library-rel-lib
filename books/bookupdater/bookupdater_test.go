package bookupdater_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/rafaelhl/library-rel-lib/books"
	"github.com/rafaelhl/library-rel-lib/books/bookupdater"
	"github.com/rafaelhl/library-rel-lib/books/bookupdater/mocks"
)

var book = books.Book{
	ID:          1,
	Title:       "Livro de Teste",
	Description: "Esse livro é de teste",
	Author:      "Rafael Holanda",
	Edition:     1,
	BookShelf: books.Shelf{
		ID: 1,
	},
}

func TestUpdater_UpdateBook(t *testing.T) {
	engine := new(mocks.UpdaterEngine)
	engine.On("UpdateBook", mock.Anything, book).Return(nil)
	updater := bookupdater.New(engine)
	err := updater.UpdateBook(context.Background(), book)
	assert.NoError(t, err)
}

func TestUpdater_UpdateBook_Error(t *testing.T) {
	engine := new(mocks.UpdaterEngine)
	engine.On("UpdateBook", mock.Anything, book).Return(errors.New("unexpected error!"))
	updater := bookupdater.New(engine)
	err := updater.UpdateBook(context.Background(), book)
	assert.Error(t, err)
}
