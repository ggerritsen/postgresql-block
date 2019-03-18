package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // justifying comment for golint
)

type repository struct {
	db *sql.DB
}

// Record is the object stored in this repository
type Record struct {
	name string
}

// NewRepositoryWithDb connects to database and uses that connection to create a repository
func NewRepositoryWithDb(host, user, password, dbname string, port int) (*repository, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return NewRepository(db), nil
}

// NewRepository creates a repository based on a db connection
func NewRepository(db *sql.DB) *repository {
	return &repository{db: db}
}

// Close should be called to properly close the repository's db connection
func (r *repository) Close() error {
	return r.db.Close()
}

// QueryByID will return the record with the provided id
func (r *repository) QueryByID(id int) (*Record, error) {
	q := "SELECT name FROM records WHERE ID = $1;"
	row := r.db.QueryRow(q, id)

	var name string
	if err := row.Scan(&name); err != nil {
		return nil, err
	}

	return &Record{name: name}, nil
}

// CreateTable will create the table in the db that is backing this repository
func (r *repository) CreateTable() error {
	q := "SELECT 1 FROM records;"
	_, err := r.db.Exec(q)
	if err == nil {
		// table already exists, nothing to do here
		return nil
	}

	q = "CREATE TABLE records (id SERIAL PRIMARY KEY, name text);"
	_, err = r.db.Exec(q)
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a record with the specified name and returns the id of the newly inserted record
func (r *repository) Insert(name string) (int, error) {
	q := "INSERT INTO records (name) VALUES ($1) RETURNING id;"

	var id int
	if err := r.db.QueryRow(q, name).Scan(&id); err != nil {
		return -1, err
	}

	return id, nil
}

// Update updates the name of record with the specified id
func (r *repository) Update(id int, updatedName string) error {
	q := "UPDATE records SET name = $1 WHERE id = $2;"
	_, err := r.db.Exec(q, updatedName, id)
	if err != nil {
		return err
	}
	return nil
}
