package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Task Manager API!")
	})

	port := "8080"
	log.Printf("Server is running on port %s\n", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}