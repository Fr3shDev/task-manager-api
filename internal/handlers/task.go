package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
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

// TaskDetailHandler handles GET, PUT, and DELETE requests for an individual task. 
// It is mapped to endpoints like "/tasks/1"
func TaskDetailHandler(w http.ResponseWriter, r *http.Request) {
	// Split the URL path to extract the task ID.
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		http.NotFound(w, r)
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	// Find the task in the in-memory slice.
	index := -1
	for i, t := range tasks {
		if t.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		// Return the specific task.
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks[index])
	case http.MethodPut:
		// Update the task.
		var updatedTask Task
		if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		updatedTask.ID = id // Ensure the task ID does not change.
		tasks[index] = updatedTask
		log.Printf("Task updated: %v\n", updatedTask)
		json.NewEncoder(w).Encode(updatedTask)
	case http.MethodDelete:
		// Delete the task by removing it from the slice.
		tasks = append(tasks[:index], tasks[index+1:]...)
		log.Printf("Task with ID %d deleted\n", id)
		w.WriteHeader(http.StatusNoContent)
	default: 
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
