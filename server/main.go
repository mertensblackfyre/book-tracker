package main

import (
	"context"
	"encoding/json"
	"fmt"
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

	// Enable all methods by default

	pkg.InitSessions()
	// r.Use(handlers.Authenticate)
	r.Use(middleware.Logger)

	// Auth
	r.Get("/auth/google", pkg.GoogleLogin)
	r.Get("/auth/callback", pkg.GoogleCallBack)

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
	r.Get("/mybooks/{status}-{user_id}", func(w http.ResponseWriter, r *http.Request) {
		status := chi.URLParam(r, "status")
		user_id := chi.URLParam(r, "user_id")
		var data pkg.Book

		err := crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
			data = pkg.FilterBooks(context.Background(), tx, status, user_id)
			return nil
		})

		if err != nil {
			log.Fatalln(err)
		}

		if err != nil {
			log.Fatalln(err)
		}

		// Encode struct to JSON
		d, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(d)
	})

	r.Get("/mybooks", func(w http.ResponseWriter, r *http.Request) {
		err := crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
			return pkg.GetAllBooks(conn)
		})

		if err != nil {
			log.Fatalln(err)
		}
	})

	r.HandleFunc("/add-book", func(w http.ResponseWriter, r *http.Request) {

		data := []byte(`{
  			"title": "The Great Gatsby",
  			"author": "F. Scott Fitzgerald",
  			"pages": "224",
  			"picture": "https://example.com/great_gatsby.jpg", 
  			"price": 10.99,
  			"status": "published",
  			"user_id": "892076584733999105"
			}`)

		err := crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
			return pkg.AddBook(tx, data)
		})

		if err != nil {
			log.Fatalln(err)
		}

		fmt.Fprintf(w, string(data), 200)
	})

	r.HandleFunc("/delete-book/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		err := crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
			return pkg.DeleteBook(context.Background(), tx, id)
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
