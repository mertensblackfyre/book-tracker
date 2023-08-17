package main

import (
	"context"
	"log"
	"net/http"

	crdbpgx "github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgxv5"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	pgx "github.com/jackc/pgx/v5"
	"github.com/server/pkg"
)

func main() {
	// Read in connection string

	conn := pkg.DBConfig()
	r := chi.NewRouter()

	pkg.InitSessions()
	// r.Use(handlers.Authenticate)
	r.Use(middleware.Logger)

	// Auth
	r.Get("/auth/google", pkg.GoogleLogin)
	r.Get("/auth/callback", func(w http.ResponseWriter, r *http.Request) {
		data := pkg.GoogleCallBack(w, r)

		err := crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
			return pkg.AddUser(tx, data)
		})

		if err != nil {
			return
		}
	})
	r.Get("/logout", func(w http.ResponseWriter, r *http.Request) {
		pkg.Logout(w, r)
	})

	// Users
	r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		err := crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
			return pkg.PrintAllUsers(conn)
		})

		if err != nil {
			log.Fatalln(err)
		}
	})

	// Books
	r.Get("/books", func(w http.ResponseWriter, r *http.Request) {
		err := crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
			return pkg.GetAllBooks(conn)
		})

		if err != nil {
			log.Fatalln(err)
		}
	})

	// Set up table
	// err := crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
	// 	return pkg.CreateTables(context.Background(), tx)
	// })

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	http.ListenAndServe(":5000", pkg.Manager.LoadAndSave(r))
	defer conn.Close(context.Background())

}
