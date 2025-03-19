package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// Task represents a simple task model
type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

// In-memory storage tasks.
var tasks []Task
var nextTaskID = 1

// TasksHandler handles requests for creating a task and listing all tasks.
// It is mapped to the "/tasks" endpoint.
func TasksHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the URL is exactly "/tasks" to avoid ambiguity.
	if r.URL.Path != "/tasks" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		// Return the tasks as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
	case http.MethodPost:
		// Decode the task from the request
		var newTask Task
		if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Basic validation: task must have a title
		if newTask.Title == "" {
			http.Error(w, "Title is required", http.StatusBadRequest)
			return
		}

		// Simulate saving the task
		newTask.ID = nextTaskID
		nextTaskID++
		tasks = append(tasks, newTask)

		log.Printf("Task created: %+v\n", newTask)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newTask)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
