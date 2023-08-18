package pkg

import (
	"context"
	"encoding/json"
	"log"

	pgx "github.com/jackc/pgx/v5"
)

func AddBook(tx pgx.Tx, data []byte) error {
	var b Book

	err := json.Unmarshal([]byte(data), &b)
	if err != nil {
		log.Println(err)
		return err
	}

	if _, err := tx.Exec(context.Background(),
		"INSERT INTO books (title ,author ,status,pages,price,picture,user_id,created_at) VALUES ($1, $2, $3, $4,$5,$6,$7, NOW())", b.Title, b.Author, b.Status, b.Pages, b.Prices, b.Picture, b.UserID); err != nil {
		log.Println(err)
		return err
	}

	log.Printf("%s has been added", b.Title)

	return nil
}

func GetAllBooks(conn *pgx.Conn) error {
	rows, err := conn.Query(context.Background(), "SELECT id, title, author,user_id FROM books;")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		var id int
		var user_id string
		var author string
		var title string
		if err := rows.Scan(&id, &title, &author, &user_id); err != nil {
			log.Fatal(err)
		}
		log.Printf("%d: %s\n -> %s", id, title, user_id)
	}

	return nil
}

func FilterBooks(conn *pgx.Conn, ctx context.Context, tx pgx.Tx, status string, user_id int) []Book {

	var books []Book
	rows, err := conn.Query(ctx, "SELECT id, title, author, status ,picture, price FROM books WHERE status = $1 AND user_id = $2", status, user_id)

	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Status, &book.Picture, &book.Prices); err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}

	log.Printf("Found %d books", len(books))

	return books
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

func UpdateBookStatus() {}

func GetUsersBook() {}
