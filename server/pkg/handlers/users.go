package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	pgx "github.com/jackc/pgx/v5"
)

func AddUser(ctx context.Context, tx pgx.Tx, data []byte) error {
	var user Users

	err := json.Unmarshal([]byte(data), &user)
	if err != nil {
		log.Fatalln(err)
	}

	// Convert UUID to string
	// id := uuid.New().String()

	// Use ExecContext instead of Exec
	if _, err := tx.Exec(ctx,
		"INSERT INTO users (email ,name ,picture,verified,created_at) VALUES ($1, $2, $3, $4, NOW())", user.Email, user.Name, user.Picture, user.VerifiedEmail); err != nil {
		fmt.Println("Hello World")
		log.Fatalln(err)
		return err
	}

	log.Println("Added users")

	return nil
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
