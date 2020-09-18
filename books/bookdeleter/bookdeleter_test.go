package bookdeleter_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/rafaelhl/library-rel-lib/books"
	"github.com/rafaelhl/library-rel-lib/books/bookdeleter"
	"github.com/rafaelhl/library-rel-lib/books/bookdeleter/mocks"
)

func TestBookDeleter_DeleteBook(t *testing.T) {
	finder := new(mocks.Finder)
	book := books.Book{
		ID:          1,
		Title:       "Deleting a book",
		Author:      "Test",
		Description: "Book of test delete",
		Edition:     1,
		ShelfID:     1,
	}
	finder.On("FindBookByID", mock.Anything, 1).Return(book, nil)
	deleter := new(mocks.Deleter)
	deleter.On("DeleteBook", mock.Anything, book).Return(nil)
	d := bookdeleter.New(finder, deleter)

	err := d.DeleteBook(context.Background(), 1)
	assert.NoError(t, err)
}
