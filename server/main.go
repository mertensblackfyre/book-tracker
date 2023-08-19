package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	crdbpgx "github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgxv5"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	pgx "github.com/jackc/pgx/v5"
	"github.com/server/pkg"
)

func JSONStruct(file string) []pkg.Book {
	// Open JSON file
	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}

	// Read opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var b []pkg.Book
	err = json.Unmarshal(byteValue, &b)
	if err != nil {
		fmt.Println(err)
	}
	return b
}

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
		user_id, err := strconv.Atoi(chi.URLParam(r, "user_id"))
		if err != nil {
			log.Println(err)
		}
		var data []pkg.Book

		err = crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
			data = pkg.FilterBooks(conn, context.Background(), tx, status, user_id)
			return nil
		})

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

	// Books
	r.Get("/mybooks/{user_id}", func(w http.ResponseWriter, r *http.Request) {
		user_id, err := strconv.Atoi(chi.URLParam(r, "user_id"))
		if err != nil {
			log.Println(err)
		}
		var data []pkg.Book

		err = crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
			data = pkg.GetUsersBooks(conn, context.Background(), tx, user_id)
			return nil
		})

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

	r.Get("/books", func(w http.ResponseWriter, r *http.Request) {
		err := crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
			return pkg.GetAllBooks(conn)
		})

		if err != nil {
			log.Fatalln(err)
		}
	})

	r.HandleFunc("/add-book", func(w http.ResponseWriter, r *http.Request) {

		data := JSONStruct("MOCK_DATA.json")

		for i := 0; i < 50; i++ {
			d, err := json.Marshal(data[i])
			err = crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
				return pkg.AddBook(tx, d)
			})

			if err != nil {
				log.Fatalln(err)
			}

			fmt.Fprintf(w, string(d), 200)

		}
	})

	r.HandleFunc("/update/{id}", func(w http.ResponseWriter, r *http.Request) {

		status := chi.URLParam(r, "status")
		book_id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			log.Println(err)
		}

		err = crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
			return pkg.UpdateBookStatus(context.Background(), tx, book_id, status)
		})

		if err != nil {
			log.Fatalln(err)
		}

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
