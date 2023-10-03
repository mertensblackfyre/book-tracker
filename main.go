package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
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

	r.Use(pkg.Authenticate)
	r.Use(middleware.Logger)

	tmpl := template.Must(template.ParseGlob("static/templates/*"))

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "index.html", nil)
	})

	r.Get("/mybooks", func(w http.ResponseWriter, r *http.Request) {
		data := r.Context().Value("data").(string)
		if data == "" {
			return
		}

		books, err := q.GetUsersBooks(data)

		if err != nil {
			log.Println(err)
			return
		}

		tmpl.ExecuteTemplate(w, "dashboard.html", books)
	})

	r.Get("/addbook", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "addbook.html", nil)
	})

	// Public routes
	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "login.html", nil)
	})

	// Auth
	r.Get("/auth/google", pkg.GoogleLogin)
	r.Get("/auth/callback", pkg.GoogleCallBack)
	r.Get("/logout", pkg.Logout)

	// Users
	//r.Get("/all/users", pkg.Han(q.AllUsers))

	// Books
	r.Post("/add-book", q.AddBook)
	r.Get("/book/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			log.Println(err)
		}

		book, err := q.GetBook(id)

		if err != nil {
			fmt.Println(err)
		}

		tmpl.ExecuteTemplate(w, "details.html", book)

	})

	r.Get("/get-{status}-{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			log.Println(err)
		}
		q.FilterBooks(chi.URLParam(r, "status"), id)
	})

	r.Delete("/delete-{id}", func(w http.ResponseWriter, r *http.Request) {
		q.DeleteBook(chi.URLParam(r, "id"))
	})

	r.Put("/change-{status}-{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		status := chi.URLParam(r, "status")

		if err != nil {
			log.Println(err)
		}

		q.UpdateBookStatus(id, status)

	})

	http.ListenAndServe(":5000", r)
}
