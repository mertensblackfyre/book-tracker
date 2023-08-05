package main

// b2pa3cQlZImdKNdj

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/server/pkg/handlers"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

    handlers.CreateTables(context.Background(),tx.)
	http.ListenAndServe(":5000", r)

}
