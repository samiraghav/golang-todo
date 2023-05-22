package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"golang-todo/backend/controllers"
	"golang-todo/backend/db"
	"golang-todo/backend/routes"
)

func main() {
	// Initialize the database connection
	err := db.InitDB()
	if err != nil {
		panic(err)
	}

	// Create a new router
	router := mux.NewRouter()

	// Register the routes
	router.HandleFunc("/", routes.HomeHandler)
	router.HandleFunc("/todo", controllers.CreateTodoHandler).Methods("POST")
	router.HandleFunc("/todo/{id}", controllers.UpdateTodoHandler).Methods("PUT")
	router.HandleFunc("/todo/{id}", controllers.DeleteTodoHandler).Methods("DELETE")
	router.HandleFunc("/todo", controllers.FetchTodosHandler).Methods("GET")

	// Serve static files
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Start the server
	http.ListenAndServe(":9000", router)
}
