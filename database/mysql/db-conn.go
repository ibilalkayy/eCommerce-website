package mysql

// Importing the libraries
import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ibilalkayy/eCommerce/middleware"
)

// It will build the connection by taking the user, password, address, db from mysql database and
// using it in the sql.Open() and return it.
func Connection() (db *sql.DB) {
	db_user := middleware.LoadEnvVariable("DB_USER")
	db_password := middleware.LoadEnvVariable("DB_PASSWORD")
	db_address := middleware.LoadEnvVariable("DB_ADDRESS")
	db_db := middleware.LoadEnvVariable("DB_DB")
	s := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", db_user, db_password, db_address, db_db)
	db, err := sql.Open("mysql", s)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// It will access the db.SQL file and the queries that are written there.
// It will split the query, prepare and then execute that query to be return at the end.
func CreateTable(file string, number int) (db *sql.DB) {
	db = Connection()
	query, err := ioutil.ReadFile("database/mysql/" + file)
	if err != nil {
		log.Fatal(err)
	}
	requests := strings.Split(string(query), ";")[number]

	stmt, err := db.Prepare(requests)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}
	return db
}
