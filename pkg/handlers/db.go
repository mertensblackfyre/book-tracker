package handlers

import (
	"context"
	"log"

	pgx "github.com/jackc/pgx/v5"
)

func DBConfig() *pgx.Conn {

	config, err := pgx.ParseConfig(GetEnv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	config.RuntimeParams["application_name"] = "$ docs_simplecrud_gopgx"
	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}

	return conn

}
func CreateTables(ctx context.Context, tx pgx.Tx) error {
	// Dropping existing table if it exists
	log.Println("Drop existing accounts table if necessary.")
	if _, err := tx.Exec(ctx, "DROP TABLE IF EXISTS books"); err != nil {
		return err
	}

	log.Println("Drop existing accounts table if necessary.")
	if _, err := tx.Exec(ctx, "DROP TABLE IF EXISTS users"); err != nil {
		return err
	}

	// Create the users table
	log.Println("Creating users table...")
	if _, err := tx.Exec(ctx,
		`CREATE TABLE users (
  		id SERIAL PRIMARY KEY,
  		email TEXT,
  		name TEXT,
  		picture TEXT,
  		verified_email BOOLEAN,
  		created_at TIMESTAMP
);
`); err != nil {
		return err
	}

	log.Println("Creating books table...")
	if _, err := tx.Exec(ctx,
		`CREATE TABLE books (
  id SERIAL PRIMARY KEY,
  title TEXT,
  author TEXT,
  user_id INTEGER,
  status TEXT,
  price REAL,
  created_at TIMESTAMP
);

`); err != nil {
		return err
	}

	log.Println("Creating foreign key...")
	if _, err := tx.Exec(ctx,
		`ALTER TABLE "books" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");`); err != nil {
		return err
	}
	return nil
}


// func transferFunds(ctx context.Context, tx pgx.Tx, from uuid.UUID, to uuid.UUID, amount int) error {
//     // Read the balance.
//     var fromBalance int
//     if err := tx.QueryRow(ctx,
//         "SELECT balance FROM accounts WHERE id = $1", from).Scan(&fromBalance); err != nil {
//         return err
//     }

//     if fromBalance < amount {
//         log.Println("insufficient funds")
//     }

//     // Perform the transfer.
//     log.Printf("Transferring funds from account with ID %s to account with ID %s...", from, to)
//     if _, err := tx.Exec(ctx,
//         "UPDATE accounts SET balance = balance - $1 WHERE id = $2", amount, from); err != nil {
//         return err
//     }
//     if _, err := tx.Exec(ctx,
//         "UPDATE accounts SET balance = balance + $1 WHERE id = $2", amount, to); err != nil {
//         return err
//     }
//     return nil
// }

// func deleteRows(ctx context.Context, tx pgx.Tx, one uuid.UUID, two uuid.UUID) error {
//     // Delete two rows into the "accounts" table.
//     log.Printf("Deleting rows with IDs %s and %s...", one, two)
//     if _, err := tx.Exec(ctx,
//         "DELETE FROM accounts WHERE id IN ($1, $2)", one, two); err != nil {
//         return err
//     }
//     return nil
// }