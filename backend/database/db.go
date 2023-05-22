package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var database *sql.DB

const (
	dbDriver   = "mysql"
	dbUsername = "root"
	dbPassword = "Samir@2002"
	dbHost     = "localhost"
	dbPort     = "3306"
	dbName     = "todo_app"
	TableName  = "todos"
)

func InitDB() error {
	// Initialize the database connection
	dbConn, err := sql.Open(dbDriver, dbUsername+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?parseTime=true")
	if err != nil {
		return err
	}
	database = dbConn
	return nil
}

func GetDB() *sql.DB {
	return database
}
