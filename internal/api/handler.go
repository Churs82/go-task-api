package api

import (
    "net/http"
    "github.com/gorilla/mux"
    "go-task-api/internal/task"
    "encoding/json"
)

type Handler struct {
    TaskManager *task.TaskManager
}

func NewHandler(tm *task.TaskManager) *Handler {
    return &Handler{TaskManager: tm}
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
    // Implementation for creating a task
}

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
    // Implementation for deleting a task
}

func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
    // Implementation for retrieving a task
}

func (h *Handler) GetTaskStatus(w http.ResponseWriter, r *http.Request) {
    // Implementation for getting task status
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
    r.HandleFunc("/tasks", h.CreateTask).Methods("POST")
    r.HandleFunc("/tasks/{id}", h.DeleteTask).Methods("DELETE")
    r.HandleFunc("/tasks/{id}", h.GetTask).Methods("GET")
    r.HandleFunc("/tasks/{id}/status", h.GetTaskStatus).Methods("GET")
}