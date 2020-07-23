package bookfind_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/rafaelhl/library-rel-lib/books"
	"github.com/rafaelhl/library-rel-lib/books/handler/bookfind"
	"github.com/rafaelhl/library-rel-lib/books/handler/bookfind/mocks"
)

var (
	body, _ = ioutil.ReadFile("mocks/book.json")
	book    = books.Book{
		ID:          1,
		Title:       "Livro de Teste",
		Description: "Esse livro é de teste",
		Author:      "Rafael Holanda",
		Edition:     1,
		BookShelf: books.Shelf{
			ID: 1,
		},
	}
)

func TestHandler_ServeHTTP(t *testing.T) {
	finder := new(mocks.BookFinder)
	finder.On("FindBook", mock.Anything, book.ID).Return(book, nil)
	handler := bookfind.NewHandler(finder)

	mux := chi.NewMux()
	mux.Method(http.MethodGet, "/books/{bookID}", handler)
	recorder := httptest.NewRecorder()
	mux.ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, fmt.Sprintf("/books/%d", 1), nil))

	finder.AssertExpectations(t)
	
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.JSONEq(t, string(body), recorder.Body.String())
}
