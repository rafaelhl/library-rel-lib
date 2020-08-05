package bookupdate_test

import (
	"errors"
	"fmt"
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
	"github.com/rafaelhl/library-rel-lib/books/handler/bookupdate"
	"github.com/rafaelhl/library-rel-lib/books/handler/bookupdate/mocks"
)

var (
	body, _ = ioutil.ReadFile("mocks/book.json")
	bookID  = 1
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
	updater := new(mocks.BookUpdater)
	b := book
	b.ID = bookID
	updater.On("UpdateBook", mock.Anything, b).Return(nil)

	path := fmt.Sprintf("/books/%d", bookID)
	recorder := sendRequest(updater, path, strings.NewReader(string(body)))

	assert.Equal(t, http.StatusOK, recorder.Code)
	updater.AssertExpectations(t)
}

func TestHandler_ServeHTTP_InvalidPayload(t *testing.T) {
	updater := new(mocks.BookUpdater)
	path := fmt.Sprintf("/books/%d", bookID)
	recorder := sendRequest(updater, path, strings.NewReader("{"))

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	updater.AssertExpectations(t)
}

func TestHandler_ServeHTTP_InvalidParam(t *testing.T) {
	updater := new(mocks.BookUpdater)
	path := fmt.Sprintf("/books/%s", "a")
	recorder := sendRequest(updater, path, strings.NewReader(string(body)))

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	updater.AssertExpectations(t)
}

func TestHandler_ServeHTTP_UpdateWithError(t *testing.T) {
	updater := new(mocks.BookUpdater)
	b := book
	b.ID = bookID
	updater.On("UpdateBook", mock.Anything, b).Return(errors.New("unexpected error!"))

	path := fmt.Sprintf("/books/%d", bookID)
	recorder := sendRequest(updater, path, strings.NewReader(string(body)))

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	updater.AssertExpectations(t)
}

func sendRequest(updater *mocks.BookUpdater, path string, body io.Reader) *httptest.ResponseRecorder {
	mux := chi.NewMux()
	handler := bookupdate.NewHandler(updater)
	mux.Method(http.MethodPut, "/books/{bookID}", handler)

	recorder := httptest.NewRecorder()
	mux.ServeHTTP(recorder, httptest.NewRequest(http.MethodPut, path, body))
	return recorder
}
