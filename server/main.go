package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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
	r := chi.NewRouter()

	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal(err)
	}

	q := pkg.NewDB(db)
	q.Drop()
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

	r.Get("/all/users", pkg.Han(q.AllUsers))

	http.ListenAndServe(":5000", pkg.Manager.LoadAndSave(r))

}
