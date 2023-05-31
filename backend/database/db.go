package db

import (
	"database/sql"
	"time"

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

	// Create the schema if it doesn't exist
	err = createSchema()
	if err != nil {
		return err
	}

	// Connect to the specific database
	err = database.Ping()
	if err != nil {
		return err
	}
	database.Exec("USE " + dbName)

	// Create the table if it doesn't exist
	err = createTable()
	if err != nil {
		return err
	}

	return nil
}

func GetDB() *sql.DB {
	return database
}

func createSchema() error {
	// Open a connection to the MySQL server without selecting a database
	conn, err := sql.Open(dbDriver, dbUsername+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/?parseTime=true")
	if err != nil {
		return err
	}
	defer conn.Close()

	// Create the schema if it doesn't exist
	_, err = conn.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		return err
	}

	return nil
}

func createTable() error {
	// Create the SQL query to create the table
	query := `CREATE TABLE IF NOT EXISTS ` + TableName + ` (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(255),
		completed BOOLEAN,
		created_at DATETIME,
		updated_at DATETIME
	)`

	// Execute the query
	_, err := database.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func AddTodoTask(title string, completed bool) (int64, error) {
	// Prepare the SQL query for inserting a new todo task
	query := "INSERT INTO " + TableName + " (title, completed, created_at) VALUES (?, ?, ?)"

	// Get the current time in Indian Standard Time (IST)
	now := time.Now().UTC().Add(time.Hour * 5).Add(time.Minute * 30) // Adding 5 hours and 30 minutes for IST offset

	// Execute the query and retrieve the inserted ID
	result, err := database.Exec(query, title, completed, now)
	if err != nil {
		return 0, err
	}

	// Get the last inserted ID
	insertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return insertedID, nil
}
