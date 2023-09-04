package pkg

import (
	"context"
	"errors"
	"fmt"
	"log"

	pgx "github.com/jackc/pgx/v5"
	sqlite3 "github.com/mattn/go-sqlite3"
)

func (r *DB) AddBook() {

	b := JSONStruct("MOCK_DATA.json")

	for i := 0; i < len(b); i++ {

		fmt.Println(i)
		response, err := r.db.Exec("INSERT INTO books (title ,author ,status,pages,price,picture,user_id) VALUES (?,?,?,?,?,?,?)", b[i].Title, b[i].Author, b[i].Status, b[i].Pages, b[i].Prices, b[i].Picture, b[i].UserID)

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
		}

		log.Println(id)

	}
}

func (r *DB) GetAllBooks() {

	rows, err := r.db.Query("SELECT * FROM books")
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

func DeleteBook(ctx context.Context, tx pgx.Tx, id string) error {
	log.Printf("Deleting book with IDs %s", id)
	if _, err := tx.Exec(ctx,
		"DELETE FROM books WHERE id IN ($1)", id); err != nil {
		return err
	}

	log.Printf("Deleted book with IDs %s", id)
	return nil
}

func UpdateBookStatus(ctx context.Context, tx pgx.Tx, book_id int, status string) error {

	log.Println("Updating status")
	if _, err := tx.Exec(ctx,
		"UPDATE books SET status = $1 WHERE id = $2", status, book_id); err != nil {
		return err
	}

	log.Println("Updated status")
	return nil
}

func (r *DB) GetUsersBooks(user_id int) []Book {

	rows, err := r.db.Query("SELECT * FROM books WHERE user_id = ?", user_id)
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
