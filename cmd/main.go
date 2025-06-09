package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "go-task-api/internal/api"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/tasks", api.CreateTask).Methods("POST")
    r.HandleFunc("/tasks/{id}", api.DeleteTask).Methods("DELETE")
    r.HandleFunc("/tasks/{id}", api.GetTask).Methods("GET")
    r.HandleFunc("/tasks/{id}/status", api.GetTaskStatus).Methods("GET")

    http.Handle("/", r)

    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}