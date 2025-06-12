package main

import (
	"go-task-api/internal/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/task", handlers.CreateTaskHandler)

	http.HandleFunc("/task/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetTaskHandler(w, r)
		case http.MethodDelete:
			handlers.DeleteTaskHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/task/{id}/status", handlers.GetTaskStatusHandler)

	http.HandleFunc("/task/{id}/result", handlers.GetTaskResultHandler)

	http.HandleFunc("/tasks", handlers.GetTaskListHandler)

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
