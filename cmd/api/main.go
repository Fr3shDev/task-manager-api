package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Fr3shDev/task-manager-api/internal/handlers"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Task Manager API!")
	})
	// User endpoints.
	http.HandleFunc("/users/register", handlers.RegisterHandler)
	http.HandleFunc("/users/login", handlers.LoginHandler)

	// Task endpoints.
	http.HandleFunc("/tasks", handlers.TasksHandler)
	http.HandleFunc("/tasks/", handlers.TaskDetailHandler)

	port := "8080"
	log.Printf("Server is running on port %s\n", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}