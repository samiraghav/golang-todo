package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/samiraghav/backend/controllers/"
)

const (
	port = ":9000"
)

func main() {
	stopSignalChan := make(chan os.Signal, 1)
	signal.Notify(stopSignalChan, os.Interrupt)

	router := mux.NewRouter()
	InitializeRoutes(router)

	// Serve static files
	staticDir := "/static/"
	staticHandler := http.StripPrefix(staticDir, http.FileServer(http.Dir("./static")))
	router.PathPrefix(staticDir).Handler(staticHandler)

	srv := &http.Server{
		Addr:         port,
		Handler:      router,
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

	<-stopSignalChan
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)
	defer cancel()
	log.Println("Server gracefully stopped!")
}

func InitializeRoutes(router *mux.Router) {
	router.HandleFunc("/", homeHandler)

	todoController := controllers.NewTodoController()

	router.HandleFunc("/todo", todoController.FetchTodosHandler).Methods("GET")
	router.HandleFunc("/todo", todoController.CreateTodoHandler).Methods("POST")
	router.HandleFunc("/todo/{id}", todoController.UpdateTodoHandler).Methods("PUT")
	router.HandleFunc("/todo/{id}", todoController.DeleteTodoHandler).Methods("DELETE")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./static/index.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
