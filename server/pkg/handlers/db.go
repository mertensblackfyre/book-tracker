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
		`CREATE TABLE "users" (
  		"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  		"email" text,
  		"name" text,
		"picture" text,
		"verified" boolean,
  		"created_at" timestamp
);`); err != nil {
		return err
	}

	log.Println("Creating books table...")
	if _, err := tx.Exec(ctx,
		`CREATE TABLE "books" (
  		"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  		"title" text,
  		"author" text,
  		"user_id" UUID,
  		"status" text,
  		"price" real,
 	 	"created_at" timestamp
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

// func printBalances(conn *pgx.Conn) error {
//     rows, err := conn.Query(context.Background(), "SELECT id, balance FROM accounts")
//     if err != nil {
//         log.Fatal(err)
//     }
//     defer rows.Close()
//     for rows.Next() {
//         var id uuid.UUID
//         var balance int
//         if err := rows.Scan(&id, &balance); err != nil {
//             log.Fatal(err)
//         }
//         log.Printf("%s: %d\n", id, balance)
//     }
//     return nil
// }

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

// func main() {
//     // Read in connection string
//     config, err := pgx.ParseConfig(os.Getenv("DATABASE_URL"))
//     if err != nil {
//         log.Fatal(err)
//     }
//     config.RuntimeParams["application_name"] = "$ docs_simplecrud_gopgx"
//     conn, err := pgx.ConnectConfig(context.Background(), config)
//     if err != nil {
//         log.Fatal(err)
//     }
//     defer conn.Close(context.Background())

//     // Set up table
//     err = crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
//         return initTable(context.Background(), tx)
//     })

//     // Insert initial rows
//     var accounts [4]uuid.UUID
//     for i := 0; i < len(accounts); i++ {
//         accounts[i] = uuid.New()
//     }

//     err = crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
//         return insertRows(context.Background(), tx, accounts)
//     })
//     if err == nil {
//         log.Println("New rows created.")
//     } else {
//         log.Fatal("error: ", err)
//     }

//     // Print out the balances
//     log.Println("Initial balances:")
//     printBalances(conn)

//     // Run a transfer
//     err = crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
//         return transferFunds(context.Background(), tx, accounts[2], accounts[1], 100)
//     })
//     if err == nil {
//         log.Println("Transfer successful.")
//     } else {
//         log.Fatal("error: ", err)
//     }

//     // Print out the balances
//     log.Println("Balances after transfer:")
//     printBalances(conn)

//     // Delete rows
//     err = crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
//         return deleteRows(context.Background(), tx, accounts[0], accounts[1])
//     })
//     if err == nil {
//         log.Println("Rows deleted.")
//     } else {
//         log.Fatal("error: ", err)
//     }

//     // Print out the balances
//     log.Println("Balances after deletion:")
//     printBalances(conn)
// }
