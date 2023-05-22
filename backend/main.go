package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	controllers "github.com/samiraghav/golang-todo/backend/controllers"
	db "github.com/samiraghav/golang-todo/backend/database"
	handlers "github.com/samiraghav/golang-todo/backend/handlers"
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
	router.HandleFunc("/", handlers.HomeHandler)
	router.HandleFunc("/todo", controllers.CreateTodoHandler).Methods("POST")
	router.HandleFunc("/todo/{id}", controllers.UpdateTodoHandler).Methods("PUT")
	router.HandleFunc("/todo/{id}", controllers.DeleteTodoHandler).Methods("DELETE")
	router.HandleFunc("/todo", controllers.FetchTodosHandler).Methods("GET")

	// Serve static files
	router.PathPrefix("/frontend/").Handler(http.StripPrefix("/frontend/", http.FileServer(http.Dir("../frontend"))))

	// Start the server
	log.Println("Server started on http://localhost:9000/")
	log.Fatal(http.ListenAndServe("localhost:9000", router))
}
