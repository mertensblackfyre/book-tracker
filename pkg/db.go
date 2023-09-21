package pkg

import (
	"database/sql"
	"log"
)

type DB struct {
	db *sql.DB
}

func NewDB(db *sql.DB) *DB {
	return &DB{
		db: db,
	}
}

func (r *DB) Drop() error {

	// Drop users table
	query := "DROP TABLE IF EXISTS users;"

	// Drop books table
	query1 := "DROP TABLE IF EXISTS books;"

	_, err := r.db.Exec(query)

	if err != nil {
		log.Fatalln(err)
	}

	_, err = r.db.Exec(query1)

	if err != nil {
		log.Fatalln(err)
	}
	return nil
}

func (r *DB) Migrate() error {

	// users table
	query := `CREATE TABLE IF NOT EXISTS users (
            id TEXT PRIMARY KEY, 
            email TEXT,
            name TEXT,
            picture TEXT,
            verified_email BOOLEAN,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
         );`

	// books table
	query1 := `CREATE TABLE IF NOT EXISTS books (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            title TEXT,
            author TEXT, 
            user_id TEXT,
            status TEXT,
            price REAL,
            picture TEXT,
            pages INTEGER,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP

        );`

	_, err := r.db.Exec(query)

	if err != nil {
		log.Fatalln(err)
	}

	_, err = r.db.Exec(query1)

	if err != nil {
		log.Fatalln(err)
	}

	return err
}

//ALTER TABLE "books" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
