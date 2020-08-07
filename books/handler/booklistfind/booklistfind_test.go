package booklistfind_test

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/rafaelhl/library-rel-lib/books"
	"github.com/rafaelhl/library-rel-lib/books/handler/booklistfind"
	"github.com/rafaelhl/library-rel-lib/books/handler/booklistfind/mocks"
)

var (
	expectedJSON, _ = ioutil.ReadFile("mocks/booklist.json")
	expectedBooks   = []books.Book{
		{
			ID:          1,
			Title:       "Livro de Teste",
			Description: "Esse livro é de teste",
			Author:      "Rafael Holanda",
			Edition:     1,
			BookShelf: books.Shelf{
				ID: 1,
			},
		},
		{
			ID:          2,
			Title:       "Livro de Teste 2",
			Description: "Esse livro é de teste 2",
			Author:      "Rafael Holanda",
			Edition:     1,
			BookShelf: books.Shelf{
				ID: 1,
			},
		},
	}
)

func TestHandler_ServeHTTP(t *testing.T) {
	finder := new(mocks.BookListFinder)
	finder.On("GetAllBooks", mock.Anything).Return(expectedBooks, nil)

	recorder := sendRequest(finder)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.JSONEq(t, string(expectedJSON), recorder.Body.String())
}

func TestHandler_ServeHTTP_Error(t *testing.T) {
	finder := new(mocks.BookListFinder)
	finder.On("GetAllBooks", mock.Anything).Return(nil, errors.New("unexpected error!"))

	recorder := sendRequest(finder)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
}

func sendRequest(finder *mocks.BookListFinder) *httptest.ResponseRecorder {
	handler := booklistfind.NewHandler(finder)
	mux := chi.NewMux()
	mux.Method(http.MethodGet, "/books", handler)
	recorder := httptest.NewRecorder()
	mux.ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/books", nil))
	return recorder
}
