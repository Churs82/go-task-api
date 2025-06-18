package handlers

import (
	"encoding/json"
	"go-task-api/internal/task"
	"net/http"
)

var manager = task.NewTaskManager()

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
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
	id := r.PathValue("id")

	task, ok := manager.GetTask(id)
	if !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)

}

func GetTaskListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(manager.GetTasks())
}

func GetTaskStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id := r.PathValue("id")
	t, ok := manager.GetTaskStatus(id)
	if !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(t)
}

func GetTaskResultHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id := r.PathValue("id")
	response, ok := manager.GetTaskResult(id)
	if !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	ok := manager.DeleteTask(id)
	if !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
