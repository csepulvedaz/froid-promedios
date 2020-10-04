package config

import (
	"database/sql"
	"fmt"
	"os"

	//Database driver import
	_ "github.com/go-sql-driver/mysql"
)

var (
	username = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	hostname = os.Getenv("DB_HOST")
	dbname   = os.Getenv("DB_NAME")
	db       *sql.DB
	err      error
)

// Connect db
func Connect() {
	d, err := sql.Open("mysql", confMysql(dbname))

	if err != nil {
		panic(err.Error())
	}
	db = d
}

// GetDB Returns a reference to the database
func GetDB() *sql.DB {
	return db
}

// CloseDB close the database conection
func CloseDB() error {
	return db.Close()
}

func confMysql(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}
