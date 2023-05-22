package crudOps

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	db "github.com/samiraghav/golang-todo/backend/database"
	"github.com/samiraghav/golang-todo/backend/models"
)

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request payload",
		})
		return
	}

	if todo.Title == "" {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"error": "The title field is required",
		})
		return
	}

	// Insert the todo into the database
	result, err := db.GetDB().Exec("INSERT INTO "+db.TableName+" (title, completed, created_at) VALUES (?, ?, NOW())", todo.Title, todo.Completed)
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to create todo",
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

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request payload",
		})
		return
	}

	if todo.Title == "" {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"error": "The title field is required",
		})
		return
	}

	// Update the todo in the database
	_, err = db.GetDB().Exec("UPDATE "+db.TableName+" SET title=?, completed=? WHERE id=?", todo.Title, todo.Completed, id)
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to update todo",
		})
		return
	}

	response := map[string]interface{}{
		"message": "Todo updated successfully",
	}
	writeJSONResponse(w, http.StatusOK, response)
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Delete the todo from the database
	_, err := db.GetDB().Exec("DELETE FROM "+db.TableName+" WHERE id=?", id)
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to delete todo",
		})
		return
	}

	response := map[string]interface{}{
		"message": "Todo deleted successfully",
	}
	writeJSONResponse(w, http.StatusOK, response)
}

func FetchTodosHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch all todos from the database
	rows, err := db.GetDB().Query("SELECT id, title, completed, created_at FROM " + db.TableName)
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to fetch todos",
		})
		return
	}
	defer rows.Close()

	todoList := []models.Todo{}
	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt)
		if err != nil {
			continue
		}

		todoList = append(todoList, todo)
	}

	response := map[string]interface{}{
		"data": todoList,
	}
	writeJSONResponse(w, http.StatusOK, response)
}

func writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
