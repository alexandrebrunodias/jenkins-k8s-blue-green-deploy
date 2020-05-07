package db

import (
	"database/sql"
	"fmt"
	"os"

	// pq Main lib
	_ "github.com/lib/pq"
)

// Connect connect to database store
func Connect() *sql.DB {
	return createConnection("store")
}

// CreateDatabase create database structure
func CreateDatabase() {
	connection := createConnection("postgres")

	_, error := connection.Exec("CREATE DATABASE store")

	if error != nil {
		fmt.Println(error)
	}

	connection = createConnection("store")
	_, error = connection.Exec("CREATE TABLE product (id serial primary key, name varchar(60), description varchar(255), price numeric(12,2), quantity integer)")

	if error != nil {
		fmt.Println(error)
	}

	defer connection.Close()
}

func createConnection(dbName string) *sql.DB {
	dbHost := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASS")
	user := os.Getenv("DB_USER")

	if dbHost == "" {
		dbHost = "localhost"
	}

	dbProperties := "user=" + user + " dbname=" + dbName + " password=" + password + " host=" + dbHost + " sslmode=disable"
	connection, err := sql.Open("postgres", dbProperties)

	if err != nil {
		panic(err.Error())
	}

	return connection
}
