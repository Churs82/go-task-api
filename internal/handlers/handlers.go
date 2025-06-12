package handlers

import (
	"encoding/json"
	"go-task-api/internal/task"
	"net/http"
)

var manager = task.NewTaskManager()

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		TaskType string `json:"task_type"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.TaskType == "" {
		http.Error(w, "Missing or invalid task_name", http.StatusBadRequest)
		return
	}
	task := manager.CreateTask(req.TaskType)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(manager.GetTasks())

	} else {
		task, ok := manager.GetTask(id)
		if !ok {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)
	}
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	ok := manager.DeleteTask(id)
	if !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
