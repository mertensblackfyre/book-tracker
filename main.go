package main

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"html/template"
	"log"
	"net/http"
	"strconv"
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

	//r.Use(pkg.Authenticate)
	r.Use(middleware.Logger)

    pkg.InitSessions()



	tmpl := template.Must(template.ParseGlob("static/templates/*"))

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "index.html", nil)
	})

	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "login.html", nil)
	})

	r.Get("/dashboard", func(w http.ResponseWriter, r *http.Request) {

		books := q.GetAllBooks(w, r)
		tmpl.ExecuteTemplate(w, "dashboard.html", books)
	})

	r.Get("/addbook", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "addbook.html", nil)

	})

	// Auth
	r.Get("/auth/google", pkg.GoogleLogin)
	r.Get("/auth/callback", pkg.GoogleCallBack)

	r.Get("/logout", pkg.Logout)

	// Users
	r.Get("/all/users", pkg.Han(q.AllUsers))

	// Books
	r.Post("/add-book", q.AddBook)

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
		if err != nil {
			log.Println(err)
		}
		q.UpdateBookStatus(id, chi.URLParam(r, "status"))
	})

	r.Get("/my-books-{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			log.Println(err)
		}
		q.GetUsersBooks(id)

	})

	http.ListenAndServe(":5000",pkg.SessionManager.LoadAndSave(r))
}
