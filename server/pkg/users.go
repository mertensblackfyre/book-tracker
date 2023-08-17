package pkg

import (
	"context"
	"encoding/json"
	"log"

	pgx "github.com/jackc/pgx/v5"
)

func AddUser(tx pgx.Tx, data []byte) error {
	var user Users

	err := json.Unmarshal([]byte(data), &user)
	if err != nil {
		log.Println(err)
		return err
	}

	check := FindUser(context.Background(), tx, user.Email)

	if check {
		log.Printf("%s User exists", user.Email)
		return nil
	}

	if _, err := tx.Exec(context.Background(),
		"INSERT INTO users (email ,name ,picture,verified_email,created_at) VALUES ($1, $2, $3, $4, NOW())", user.Email, user.Name, user.Picture, user.VerifiedEmail); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func FindUser(ctx context.Context, tx pgx.Tx, email string) bool {
	// Read the balance.
	var e string

	if err := tx.QueryRow(ctx,
		"SELECT email FROM users WHERE email = $1", email).Scan(&e); err != nil {
	}

	if len(e) != 0 {
		return true
	}

	return false
}

func PrintAllUsers(conn *pgx.Conn) error {
	rows, err := conn.Query(context.Background(), "SELECT id, name, email FROM users;")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		var id int
		var fullname string
		var email string
		if err := rows.Scan(&id, &fullname, &email); err != nil {
			log.Fatal(err)
		}
		log.Printf("%d: %s\n", id, fullname)
	}

	return nil
}
