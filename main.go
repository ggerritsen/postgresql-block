package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Printf("Start of demonstration.\n")

	repo, err := NewRepositoryWithDb("localhost", "postgres", "your-password", "template1", 5432)
	if err != nil {
		log.Fatal(err)
	}
	defer repo.Close()
	fmt.Printf("Connected to database\n")

	if err := repo.CreateTable(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Table created\n")

	recordName := "testRecord"
	id, err := repo.Insert(recordName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted record with id %d\n", id)

	record, err := repo.QueryByID(id)
	if err != nil {
		log.Fatal(err)
	}
	if record.name != recordName {
		log.Fatal("Record in DB (%q) doesn't match expectation (%q)\n", record.name, recordName)
	}
	fmt.Printf("Successfully queried record with id %d: %+v\n", id, record)

	updatedName := fmt.Sprintf("record-%d", time.Now().Unix())
	if err := repo.Update(id, updatedName); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated record with id %d\n", id)

	record, err = repo.QueryByID(id)
	if err != nil {
		log.Fatal(err)
	}
	if record.name != updatedName {
		log.Fatal("Record in DB (%q) doesn't match expectation (%q)\n", record.name, updatedName)
	}
	fmt.Printf("Successfully queried record with id %d: %+v\n", id, record)

	fmt.Printf("End of demonstration.\n")
}
