package main

import (
	"context"
	"log"
	"net/http"

	crdbpgx "github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgxv5"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	pgx "github.com/jackc/pgx/v5"
	"github.com/server/pkg/handlers"
)

func main() {
	// Read in connection string

	conn := handlers.DBConfig()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Get("/auth/google", handlers.GoogleLogin)
	r.Get("/auth/callback", handlers.GoogleCallBack)
	r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		err := crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
			return handlers.PrintAllUsers(conn)
		})

		if err != nil {
			log.Fatalln(err)
		}
	})

	// Set up table
	err := crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		return handlers.CreateTables(context.Background(), tx)
	})

	if err != nil {
		log.Fatalln(err)
	}

	http.ListenAndServe(":5000", r)

	defer conn.Close(context.Background())

}
