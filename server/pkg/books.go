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

func FilterBooks(ctx context.Context, tx pgx.Tx, status string, user_id string) Book {
	var b Book

	if err := tx.QueryRow(ctx,
		"SELECT * FROM books WHERE status = $1 AND user_id = $2", status, user_id).Scan(&b); err != nil {
		log.Fatalln("Book not found")
	}

	return b
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
