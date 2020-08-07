package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Fs02/rel"
	"github.com/Fs02/rel/adapter/mysql"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	"github.com/rafaelhl/library-rel-lib/books/bookfinder"
	"github.com/rafaelhl/library-rel-lib/books/booklistfinder"
	"github.com/rafaelhl/library-rel-lib/books/booksinserter"
	"github.com/rafaelhl/library-rel-lib/books/bookupdater"
	"github.com/rafaelhl/library-rel-lib/books/handler/bookfind"
	"github.com/rafaelhl/library-rel-lib/books/handler/bookinsert"
	"github.com/rafaelhl/library-rel-lib/books/handler/booklistfind"
	"github.com/rafaelhl/library-rel-lib/books/handler/bookupdate"
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
	repository := repository.New(repo)

	router.Method(http.MethodPost, "/books", bookinsert.NewHandler(booksinserter.New(repository, repository)))
	router.Method(http.MethodGet, "/books/{bookID}", bookfind.NewHandler(bookfinder.New(repository)))
	router.Method(http.MethodPut, "/books/{bookID}", bookupdate.NewHandler(bookupdater.New(repository)))
	router.Method(http.MethodGet, "/books", booklistfind.NewHandler(booklistfinder.New(repository)))

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
