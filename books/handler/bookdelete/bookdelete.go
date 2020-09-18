//go:generate mockery -name=BookDeleter

package bookdelete

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type (
	BookDeleter interface {
		DeleteBook(ctx context.Context, id int) error
	}

	handler struct {
		deleter BookDeleter
	}
)

func NewHandler(deleter BookDeleter) handler {
	return handler{
		deleter: deleter,
	}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "bookID")
	bookID, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.deleter.DeleteBook(r.Context(), bookID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
