package config

import (
	"database/sql"
	"fmt"

	//Database driver import
	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "admin"
	password = "Tesla2119"
	hostname = "supermarket-db.clv7vchexipa.us-east-1.rds.amazonaws.com:3306"
	dbname   = "froid_historia_academica_db"
	// user     = os.Getenv("DB_USER")
	// password = os.Getenv("DB_PASSWORD")
	// enpoint  = os.Getenv("DB_ENDPOINT")
	// port     = os.Getenv("DB_PORT")
	// name     = os.Getenv("DB_NAME")
)

var (
	db  *sql.DB
	err error
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
