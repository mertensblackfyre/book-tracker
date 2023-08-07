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
	config, err := pgx.ParseConfig(handlers.GetEnv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	config.RuntimeParams["application_name"] = "$ docs_simplecrud_gopgx"
	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close(context.Background())

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Get("/auth/google", handlers.GoogleLogin)
	r.Get("/auth/callback", handlers.GoogleCallBack)

	// Set up table
	err = crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		return handlers.CreateTables(context.Background(), tx)
	})

	if err != nil {
		log.Fatalln(err)
	}
	http.ListenAndServe(":5000", r)

}
