package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type TaskFunc func() (string, error)

var taskRegistry = map[string]TaskFunc{}

// RegisterTask allows tasks to register themselves by name
func RegisterTask(name string, fn TaskFunc) {
	taskRegistry[name] = fn
}

type TaskStatus string

const (
	StatusPending   TaskStatus = "pending"
	StatusRunning   TaskStatus = "running"
	StatusFinished  TaskStatus = "finished"
	StatusFailed    TaskStatus = "failed"
	StatusCancelled TaskStatus = "cancelled"
)

type Task struct {
	ID        string     `json:"id"`
	Status    TaskStatus `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	Duration  float64    `json:"duration"`
	Result    string     `json:"result,omitempty"`
	Error     string     `json:"error,omitempty"`
	TaskName  string     `json:"task_name,omitempty"`
}

type TaskManager struct {
	mu    sync.RWMutex
	tasks map[string]*Task
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks: make(map[string]*Task),
	}
}

func (tm *TaskManager) CreateTask(taskName string) *Task {
	id := strconv.FormatInt(time.Now().UnixNano()+rand.Int63n(1000), 36)
	task := &Task{
		ID:        id,
		Status:    StatusPending,
		CreatedAt: time.Now(),
		TaskName:  taskName,
	}
	tm.mu.Lock()
	tm.tasks[id] = task
	tm.mu.Unlock()

	go tm.runTask(task)
	return task
}

func (tm *TaskManager) runTask(task *Task) {
	tm.mu.Lock()
	task.Status = StatusRunning
	tm.mu.Unlock()

	start := time.Now()
	fn, ok := taskRegistry[task.TaskName]
	if !ok {
		tm.mu.Lock()
		task.Status = StatusFailed
		task.Error = "Task not found: " + task.TaskName
		tm.mu.Unlock()
		return
	}
	result, err := fn()

	tm.mu.Lock()
	defer tm.mu.Unlock()
	task.Duration = time.Since(start).Seconds()
	if err != nil {
		task.Status = StatusFailed
		task.Error = err.Error()
	} else {
		task.Status = StatusFinished
		task.Result = result
	}
}

func (tm *TaskManager) GetTask(id string) (*Task, bool) {
	tm.mu.RLock()
	defer tm.mu.RUnlock()
	task, ok := tm.tasks[id]
	return task, ok
}

func (tm *TaskManager) DeleteTask(id string) bool {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	_, ok := tm.tasks[id]
	if ok {
		delete(tm.tasks, id)
	}
	return ok
}

var manager = NewTaskManager()

func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		TaskName string `json:"task_name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.TaskName == "" {
		http.Error(w, "Missing or invalid task_name", http.StatusBadRequest)
		return
	}
	task := manager.CreateTask(req.TaskName)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	task, ok := manager.GetTask(id)
	if !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	ok := manager.DeleteTask(id)
	if !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			createTaskHandler(w, r)
		case http.MethodGet:
			getTaskHandler(w, r)
		case http.MethodDelete:
			deleteTaskHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
