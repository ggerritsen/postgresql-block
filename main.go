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

	r := &Record{name: "testRecord"}
	id, err := repo.Insert(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted record with id %d\n", id)

	record, err := repo.QueryByID(id)
	if err != nil {
		log.Fatal(err)
	}
	if record.name != r.name {
		log.Fatalf("Record in DB (%q) doesn't match expectation (%q)\n", record.name, r.name)
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
		log.Fatalf("Record in DB (%q) doesn't match expectation (%q)\n", record.name, updatedName)
	}
	fmt.Printf("Successfully queried record with id %d: %+v\n", id, record)

	if err := repo.Delete(id); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted record with id %d\n", id)

	record, err = repo.QueryByID(id)
	if err != nil {
		log.Fatal(err)
	}
	if record != nil {
		fmt.Printf("Expected no record with %d anymore, but it still exists: %q\n", id, record)
	}
	fmt.Printf("Verified deletion of record with id %d\n", id)

	fmt.Printf("End of demonstration.\n")
}
