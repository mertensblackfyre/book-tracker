package main

import (
	"context"
	"html/template"
	"log"
	"net/http"

	crdbpgx "github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgxv5"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	pgx "github.com/jackc/pgx/v5"
	"github.com/server/pkg"
)

var tmpl *template.Template

func main() {
	// Read in connection string

	conn := pkg.DBConfig()
	r := chi.NewRouter()

	pkg.InitSessions()
	// r.Use(handlers.Authenticate)
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl = template.Must(template.ParseFiles("./static/index.html"))
		tmpl.Execute(w, nil)
	})

	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		tmpl = template.Must(template.ParseFiles("./static/login.html"))
		tmpl.Execute(w, nil)
	})
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

	// // // Set up table
	err := crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		return pkg.CreateTables(context.Background(), tx)
	})

	if err != nil {
		log.Fatalln(err)
	}

	http.ListenAndServe(":8000", pkg.Manager.LoadAndSave(r))
	defer conn.Close(context.Background())

}
