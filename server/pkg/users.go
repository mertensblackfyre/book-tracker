package pkg

import (
	"context"
	"encoding/json"
	"log"

	pgx "github.com/jackc/pgx/v5"
)

func AddUser( tx pgx.Tx, data []byte) error {
	var user Users

	err := json.Unmarshal([]byte(data), &user)
	if err != nil {
		log.Println(err)
		return err
	}

	if _, err := tx.Exec(context.Background(),
		"INSERT INTO users (email ,name ,picture,verified_email,created_at) VALUES ($1, $2, $3, $4, NOW())", user.Email, user.Name, user.Picture, user.VerifiedEmail); err != nil {
		log.Println(err)
		return err
	}

	// Store a new key and value in the session data.
	Manager.Put(context.Background(), "name", user.Email)

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
