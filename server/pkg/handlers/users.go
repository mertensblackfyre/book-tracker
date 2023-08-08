package handlers

import (
	"context"
	"encoding/json"
	"log"

	"github.com/google/uuid"
	pgx "github.com/jackc/pgx/v5"
)

func AddUser(ctx context.Context, tx pgx.Tx, data []byte) error {
	var user User

	err := json.Unmarshal([]byte(data), &user)
	if err != nil {
		log.Fatalln(err)
	}

	if _, err := tx.Exec(ctx,
		"INSERT INTO users (id,email ,name ,picture) VALUES ($1, $2, $3, $4)", uuid.New(), user.Email, user.Name, user.Picture); err != nil {
		return err
	}

	log.Println("Added users")
	
	return nil
}

func PrintAllUsers(conn *pgx.Conn) error {
	rows, err := conn.Query(context.Background(), "SELECT * FROM users;")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		var id uuid.UUID
		var fullname string
		var email string
		if err := rows.Scan(&id, &fullname, &email); err != nil {
			log.Fatal(err)
		}
		log.Printf("%s: %s\n", id, fullname)
	}
	log.Println("Debugger")

	return nil
}
