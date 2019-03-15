package main

import (
	"fmt"
	"log"
)

func main() {
	c := &Client{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "your-password",
		Dbname:   "template1",
	}

	db, err := c.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Printf("Initialized database\n")

	s, err := Query(db, "SELECT * FROM users LIMIT 1;")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Retrieved %q from db\n", s)
}
