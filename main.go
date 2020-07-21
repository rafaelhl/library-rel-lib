package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Fs02/rel"
	"github.com/Fs02/rel/adapter/mysql"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	"github.com/rafaelhl/library-rel-lib/books/booksinserter"
	"github.com/rafaelhl/library-rel-lib/books/handler/bookinsert"
	"github.com/rafaelhl/library-rel-lib/books/repository"
)

var dsn = "root:root@(localhost:3306)/library?parseTime=true"

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		render.PlainText(w, r, "pong")
	})

	adapter, err := mysql.Open(dsn)
	if err != nil {
		panic(err)
	}
	defer adapter.Close()

	// initialize rel's repo.
	repo := rel.New(adapter)

	router.Method(http.MethodPost, "/books", bookinsert.NewHandler(createBookInserter(repo)))

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}

func createBookInserter(repo rel.Repository) booksinserter.Inserter {
	r := repository.New(repo)
	return booksinserter.New(r, r)
}
