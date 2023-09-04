package pkg

import (
	"encoding/json"
	"errors"
	"log"

	sqlite3 "github.com/mattn/go-sqlite3"
)

func (r *DB) AddUser(data string) error {
	var user Users

	err := json.Unmarshal([]byte(data), &user)

	if err != nil {
		log.Println(err)
		return err
	}

	response, err := r.db.Exec("INSERT INTO users (email ,name ,picture,verified_email) VALUES (?, ?, ?, ?)", user.Email, user.Name, user.Picture, user.VerifiedEmail)

	if err != nil {

		log.Println(err)
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			if errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
				log.Println(err)
				return errors.New("Duplicate account")
			}
		}
		return err
	}

	if err != nil {
		return err
	}
	id, err := response.LastInsertId()
	if err != nil {
		return err
	}

	log.Println(id)

	return nil
}



func (r *DB) AllUsers() {

	rows, err := r.db.Query("SELECT * FROM users")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var all []Users
	for rows.Next() {
		var users Users
		if err := rows.Scan(&users.ID, &users.Email, &users.Name, &users.Picture, &users.VerifiedEmail, &users.Created_at); err != nil {
			log.Println(err)
		}
		all = append(all, users)
	}

	log.Println(all)
}
