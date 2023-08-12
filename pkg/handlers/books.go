package handlers

import (
	"context"
	"encoding/json"
	"log"

	pgx "github.com/jackc/pgx/v5"
)

func AddBook(ctx context.Context, tx pgx.Tx, data []byte) error {
	var b Book
	err := json.Unmarshal([]byte(data), &b)
	if err != nil {
		log.Println(err)
		return err
	}

	if _, err := tx.Exec(ctx,
		"INSERT INTO users (title ,author ,status,pages,prices,picture,user_id,created_at) VALUES ($1, $2, $3, $4,$5,$6,$7, NOW())", b.Title, b.Author, b.Status, b.Pages, b.Prices, b.Picture, b.UserID); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func GetAllBooks() {}

func FilterBooks() {}

func DeleteBook() {}

func UpdateBookStatus() {}
