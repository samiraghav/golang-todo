package main

import (
	"github.com/gorilla/mux"
	"github.com/samiraghav/backend/controllers"
)

func InitializeRoute(router *mux.Router) {
	router.HandleFunc("/", homeHandler)

	todoController := controllers.NewTodoController()

	router.HandleFunc("/todo", todoController.FetchTodosHandler).Methods("GET")
	router.HandleFunc("/todo", todoController.CreateTodoHandler).Methods("POST")
	router.HandleFunc("/todo/{id}", todoController.UpdateTodoHandler).Methods("PUT")
	router.HandleFunc("/todo/{id}", todoController.DeleteTodoHandler).Methods("DELETE")
}
