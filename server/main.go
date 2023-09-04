package main

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"html/template"
	"log"
	"net/http"

	"github.com/server/pkg"
)

func main() {
	r := chi.NewRouter()

	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal(err)
	}

	q := pkg.NewDB(db)
	//q.Drop()
	q.Migrate()

	pkg.InitSessions()
	// r.Use(handlers.Authenticate)
	r.Use(middleware.Logger)

	tmpl := template.Must(template.ParseGlob("static/*"))

	r.Use(cors.Handler(cors.Options{
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowedOrigins: []string{"http://localhost:3000"},
		//	AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "index.html", nil)
	})

	// Auth
	r.Get("/auth/google", pkg.GoogleLogin)
	r.Get("/auth/callback", pkg.GoogleCallBack)

	r.Get("/logout", func(w http.ResponseWriter, r *http.Request) {
		pkg.Logout(w, r)
	})

	// Users
	r.Get("/all/users", pkg.Han(q.AllUsers))

	// Books
	r.Get("/add/book", pkg.Han(q.AddBook))
	r.Get("/all/books", pkg.Han(q.GetAllBooks))

	http.ListenAndServe(":5000", pkg.Manager.LoadAndSave(r))
}
