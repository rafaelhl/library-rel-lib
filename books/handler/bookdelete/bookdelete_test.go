package bookdelete_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/rafaelhl/library-rel-lib/books/handler/bookdelete"
	"github.com/rafaelhl/library-rel-lib/books/handler/bookdelete/mocks"
)

func TestHandler_ServeHTTP(t *testing.T) {
	deleter := new(mocks.BookDeleter)
	deleter.On("DeleteBook", mock.Anything, 1).Return(nil)
	h := bookdelete.NewHandler(deleter)

	mux := chi.NewMux()
	mux.Method(http.MethodDelete, "/books/{bookID}", h)

	req := httptest.NewRequest(http.MethodDelete, "/books/1", nil)
	recorder := httptest.NewRecorder()
	mux.ServeHTTP(recorder, req)

	deleter.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, recorder.Code)
}
