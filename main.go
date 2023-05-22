package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB

const (
	dbDriver   = "mysql"
	dbUsername = "root"
	dbPassword = "Samir@2002"
	dbHost     = "localhost"
	dbPort     = "3306"
	dbName     = "todo_app"
	tableName  = "todos"
	port       = ":9000"
)

type (
	todo struct {
		ID        int64     `json:"id"`
		Title     string    `json:"title"`
		Completed bool      `json:"completed"`
		CreatedAt time.Time `json:"created_at"`
	}
)

func init() {
	dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?parseTime=true"
	dbConn, err := sql.Open(dbDriver, dsn)
	if err != nil {
		log.Fatal(err)
	}
	db = dbConn
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./static/index.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func createTodoHandler(w http.ResponseWriter, r *http.Request) {
	var t todo

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	if t.Title == "" {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"error": "The title field is required",
		})
		return
	}

	stmt, err := db.Prepare("INSERT INTO " + tableName + " (title, completed, created_at) VALUES (?, ?, ?)")
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to save todo: " + err.Error(),
		})
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(t.Title, false, time.Now())
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to save todo: " + err.Error(),
		})
		return
	}

	insertID, _ := result.LastInsertId()

	response := map[string]interface{}{
		"message": "Todo created successfully",
		"todo_id": insertID,
	}
	writeJSONResponse(w, http.StatusCreated, response)
}

func updateTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var t todo

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	if t.Title == "" {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"error": "The title field is required",
		})
		return
	}

	stmt, err := db.Prepare("UPDATE " + tableName + " SET title=?, completed=? WHERE id=?")
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to update todo: " + err.Error(),
		})
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(t.Title, t.Completed, id)
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to update todo: " + err.Error(),
		})
		return
	}

	response := map[string]interface{}{
		"message": "Todo updated successfully",
	}
	writeJSONResponse(w, http.StatusOK, response)
}

func fetchTodosHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, title, completed, created_at FROM " + tableName)
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to fetch todos: " + err.Error(),
		})
		return
	}
	defer rows.Close()

	todoList := []todo{}
	for rows.Next() {
		var t todo
		err := rows.Scan(&t.ID, &t.Title, &t.Completed, &t.CreatedAt)
		if err != nil {
			log.Println(err)
			continue
		}

		todoList = append(todoList, t)
	}

	response := map[string]interface{}{
		"data": todoList,
	}
	writeJSONResponse(w, http.StatusOK, response)
}

func deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	stmt, err := db.Prepare("DELETE FROM " + tableName + " WHERE id=?")
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to delete todo: " + err.Error(),
		})
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to delete todo: " + err.Error(),
		})
		return
	}

	response := map[string]interface{}{
		"message": "Todo deleted successfully",
	}
	writeJSONResponse(w, http.StatusOK, response)
}

func writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func main() {
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)

	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/todo", fetchTodosHandler).Methods("GET")
	r.HandleFunc("/todo", createTodoHandler).Methods("POST")
	r.HandleFunc("/todo/{id}", updateTodoHandler).Methods("PUT")
	r.HandleFunc("/todo/{id}", deleteTodoHandler).Methods("DELETE")

	// Serve static files
	staticDir := "/static/"
	staticHandler := http.StripPrefix(staticDir, http.FileServer(http.Dir("./static")))
	r.PathPrefix(staticDir).Handler(staticHandler)

	srv := &http.Server{
		Addr:         port,
		Handler:      r,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Println("Listening on port", port)
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	<-stopChan
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)
	defer cancel()
	log.Println("Server gracefully stopped!")
}
