package bookinsert_test

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/rafaelhl/library-rel-lib/books"
	"github.com/rafaelhl/library-rel-lib/books/handler/bookinsert"
	"github.com/rafaelhl/library-rel-lib/books/handler/bookinsert/mocks"
)

var (
	body, _ = ioutil.ReadFile("mocks/book.json")
	book    = books.Book{
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
	inserter := new(mocks.BookInserter)
	inserter.On("InsertBook", mock.Anything, book).Return(nil)

	recorder := sendRequest(inserter, strings.NewReader(string(body)))

	inserter.AssertExpectations(t)

	assert.Equal(t, http.StatusOK, recorder.Code)
	inserter.AssertExpectations(t)
}

func TestHandler_ServeHTTP_InvalidPayload(t *testing.T) {
	inserter := new(mocks.BookInserter)
	recorder := sendRequest(inserter, strings.NewReader("{"))

	inserter.AssertExpectations(t)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	inserter.AssertExpectations(t)
}

func TestHandler_ServeHTTP_InsertFail(t *testing.T) {
	inserter := new(mocks.BookInserter)
	inserter.On("InsertBook", mock.Anything, book).Return(errors.New("unexpected error!"))

	recorder := sendRequest(inserter, strings.NewReader(string(body)))

	inserter.AssertExpectations(t)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	inserter.AssertExpectations(t)
}

func sendRequest(inserter *mocks.BookInserter, body io.Reader) *httptest.ResponseRecorder {
	mux := chi.NewMux()
	handler := bookinsert.NewHandler(inserter)
	mux.Method(http.MethodPost, "/books", handler)

	recorder := httptest.NewRecorder()
	mux.ServeHTTP(recorder, httptest.NewRequest(http.MethodPost, "/books", body))
	return recorder
}
