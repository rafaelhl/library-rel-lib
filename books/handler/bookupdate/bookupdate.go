//go:generate mockery -name=BookUpdater

package bookupdate

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/rafaelhl/library-rel-lib/books"
)

type (
	BookUpdater interface {
		UpdateBook(ctx context.Context, book books.Book) error
	}

	handler struct {
		updater BookUpdater
	}
)

func NewHandler(updater BookUpdater) handler {
	return handler{
		updater: updater,
	}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var book books.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	param := chi.URLParam(r, "bookID")
	bookID, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	book.ID = bookID
	err = h.updater.UpdateBook(r.Context(), book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
