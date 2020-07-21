//go:generate mockery -name=BookInserter

package bookinsert

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/rafaelhl/library-rel-lib/books"
)

type (
	BookInserter interface {
		InsertBook(ctx context.Context, book books.Book) error
	}

	handler struct {
		inserter BookInserter
	}
)

func NewHandler(inserter BookInserter) handler {
	return handler{
		inserter: inserter,
	}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var book books.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.inserter.InsertBook(r.Context(), book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
