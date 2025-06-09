package task

import (
	"sync"
	"time"
)

type Task struct {
	TaskID      string
	Status      string
	CreationDate time.Time
	Duration    time.Duration
}

type TaskManager struct {
	tasks map[string]*Task
	mu    sync.Mutex
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks: make(map[string]*Task),
	}
}

func (tm *TaskManager) StartTask(taskID string) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	task := &Task{
		TaskID:      taskID,
		Status:      "running",
		CreationDate: time.Now(),
	}
	tm.tasks[taskID] = task
}

func (tm *TaskManager) StopTask(taskID string) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	if task, exists := tm.tasks[taskID]; exists {
		task.Status = "completed"
		task.Duration = time.Since(task.CreationDate)
	}
}

func (tm *TaskManager) GetTaskResults(taskID string) (*Task, bool) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	task, exists := tm.tasks[taskID]
	return task, exists
}