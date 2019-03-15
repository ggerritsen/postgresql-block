package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // justifying comment for golint
)

// Client provides a means to connecting to a postgresql db.
type Client struct {
	Host, User, Password, Dbname string
	Port                         int
}

// Connect connects to the postgresql db.
// Callers are expected to call Close() on the returned *sql.DB.
func (c *Client) Connect() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.Dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Query(db *sql.DB, q string) (string, error) {
	return "", nil
}

// row := db.QueryRow(q)
// switch err := row.Scan(&id, &email); err {
// case sql.ErrNoRows:
//   fmt.Println("No rows were returned!")
// case nil:
//   fmt.Println(id, email)
// default:
//   panic(err)
// }
// }
