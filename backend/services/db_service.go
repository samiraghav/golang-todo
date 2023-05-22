package services

import (
	"database/sql"
	"time"

	"github.com/your-username/backend/models"
)

type DBService struct {
	db *sql.DB
}

func NewDBService() *DBService {
	dbService := &DBService{
		db: GetDBConnection(),
	}
	return dbService
}

func (ds *DBService) CreateTodo(todo *models.Todo) (*models.Todo, error) {
	stmt, err := ds.db.Prepare("INSERT INTO " + tableName + " (title, completed, created_at) VALUES (?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(todo.Title, false, time.Now())
	if err != nil {
		return nil, err
	}

	insertID, _ := result.LastInsertId()

	todo.ID = insertID

	return todo, nil
}

func (ds *DBService) UpdateTodo(id int64, todo *models.Todo) (*models.Todo, error) {
	stmt, err := ds.db.Prepare("UPDATE " + tableName + " SET title=?, completed=? WHERE id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(todo.Title, todo.Completed, id)
	if err != nil {
		return nil, err
	}

	todo.ID = id

	return todo, nil
}

func (ds *DBService) FetchTodos() ([]models.Todo, error) {
	rows, err := ds.db.Query("SELECT id, title, completed, created_at FROM " + tableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todoList := []models.Todo{}
	for rows.Next() {
		var todoData models.Todo
		err := rows.Scan(&todoData.ID, &todoData.Title, &todoData.Completed, &todoData.CreatedAt)
		if err != nil {
			continue
		}

		todoList = append(todoList, todoData)
	}

	return todoList, nil
}

func (ds *DBService) DeleteTodo(id int64) error {
	stmt, err := ds.db.Prepare("DELETE FROM " + tableName + " WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func GetDBConnection() *sql.DB {
	dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?parseTime=true"
	dbConn, err := sql.Open(dbDriver, dsn)
	if err != nil {
		panic(err)
	}
	return dbConn
}
