package pkg

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	sqlite3 "github.com/mattn/go-sqlite3"
)

func (q *DB) AddBook(w http.ResponseWriter, r *http.Request) {

	data := r.Context().Value("data").(string)
	// Read request body
	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	// Unmarshal JSON
	var b Book
	err = json.Unmarshal(res, &b)

	if err != nil {
		log.Println(err)
		JSONWritter(w, 400, err)
	}

	response, err := q.db.Exec("INSERT INTO books (title ,author,status,pages,price,picture,user_id) VALUES (?,?,?,?,?,?,?)", b.Title, b.Author, b.Status, b.Pages, b.Prices, b.Picture, data)

	if err != nil {

		log.Println(err)
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			if errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
				log.Println(err)
			}
		}
	}

	if err != nil {
		log.Println(err)
	}

	id, err := response.LastInsertId()
	if err != nil {
		log.Println(err)
        return
	}

	log.Println(id)

}

func (q *DB) GetAllBooks(w http.ResponseWriter, r *http.Request) []Book {

	rows, err := q.db.Query("SELECT * FROM books")
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	var all []Book
	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Status, &b.Pages, &b.Prices, &b.Picture, &b.UserID, &b.Created_at); err != nil {
			log.Println(err)

		}
		all = append(all, b)
	}

	return all
}

func (r *DB) FilterBooks(status string, user_id int) []Book {
	rows, err := r.db.Query("SELECT * FROM books WHERE status = ? AND user_id = ?", status, user_id)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var all []Book
	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Status, &b.Pages, &b.Prices, &b.Picture, &b.UserID, &b.Created_at); err != nil {
			log.Println(err)
		}
		all = append(all, b)
	}

	log.Println(all)
	return all
}

func (r *DB) DeleteBook(id string) {

	res, err := r.db.Exec("DELETE FROM websites WHERE id = ?", id)
	if err != nil {
		log.Println(err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
	}

	if rowsAffected == 0 {
		log.Println(err)
	}

}

func (r *DB) UpdateBookStatus(book_id int, status string) {

	res, err := r.db.Exec("UPDATE books SET status = ?  WHERE id = ?", status, book_id)
	if err != nil {
	}

	if err != nil {
		log.Println(err)
	}

	log.Println(res)
}

func (r *DB) GetBook(id int) (Book, error) {

	row := r.db.QueryRow("SELECT * FROM books WHERE id = ?", id)

	var b Book

	if err := row.Scan(&b.ID, &b.Title, &b.Author, &b.UserID, &b.Status, &b.Prices, &b.Picture, &b.Pages, &b.Created_at); err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			log.Println(err)
		}

		log.Println(err)
	}

	return b, nil
}

func (r *DB) GetUsersBooks(user_id string) ([]Book, error) {

	rows, err := r.db.Query("SELECT * FROM books WHERE user_id = ?", user_id)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	var all []Book
	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.UserID, &b.Status, &b.Prices, &b.Picture, &b.Pages, &b.Created_at); err != nil {

			log.Println(err)
		}
		all = append(all, b)
	}

	return all, nil
}
