package main

import (
	"database/sql"
	"fmt"
)

const (
	dbUser     = "root"
	dbPassword = ""
	dbName     = "go_02"
	dbHost     = "127.0.0.1"
	dbPort     = "3306"
)

func ConnectToMySQLDB() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MySQL database successfully!")
	return db, nil
}
