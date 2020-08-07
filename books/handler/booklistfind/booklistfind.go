//go:generate mockery -name=BookListFinder

package booklistfind

import (
	"context"
	"net/http"

	"github.com/go-chi/render"

	"github.com/rafaelhl/library-rel-lib/books"
)

type (
	BookListFinder interface {
		GetAllBooks(ctx context.Context) ([]books.Book, error)
	}

	handler struct {
		finder BookListFinder
	}
)

func NewHandler(finder BookListFinder) handler {
	return handler{
		finder: finder,
	}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	allBooks, err := h.finder.GetAllBooks(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, allBooks)
}
