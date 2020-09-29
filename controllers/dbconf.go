package controllers

import (
	"database/sql"

	//Database driver import
	_ "github.com/go-sql-driver/mysql"
)

var (
	// DB The database connection
	db  *sql.DB
	err error
	// user     = os.Getenv("DB_USER")
	// password = os.Getenv("DB_PASSWORD")
	// enpoint  = os.Getenv("DB_ENDPOINT")
	// port     = os.Getenv("DB_PORT")
	// name     = os.Getenv("DB_NAME")
)

// Setup Sets up the many many app settings
func Setup() {
	// d, err := sql.Open("mysql", "admin:Tesla2119@tcp(supermarket-db.clv7vchexipa.us-east-1.rds.amazonaws.com:3306)/supermarket-db")
	d, err := sql.Open("mysql", "csepulvedaz:Camilo1015$@tcp(arquisoftlab2.ctapo3sygnnp.us-east-1.rds.amazonaws.com:3306)/supermarket_db")
	// d, err := sql.Open("mysql", user+":"+password+"@tcp("+enpoint+":"+port+")/"+name)
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
