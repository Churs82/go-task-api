package task

import (
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type TaskManager struct {
	mu    sync.RWMutex
	tasks map[string]*Task
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks: make(map[string]*Task),
	}
}

func (tm *TaskManager) CreateTask(taskType string) *Task {
	id := strconv.FormatInt(time.Now().UnixNano()+rand.Int63n(1000), 36)
	task := &Task{
		ID:        id,
		Status:    StatusPending,
		CreatedAt: time.Now(),
		TaskType:  taskType,
	}
	tm.mu.Lock()
	tm.tasks[id] = task
	tm.mu.Unlock()

	go tm.runTask(task)
	return task
}

func (tm *TaskManager) runTask(t *Task) {
	tm.mu.Lock()
	t.Status = StatusRunning
	start := time.Now()
	tm.mu.Unlock()
	// Run in parallel without blocking
	result, err := t.Run()
	tm.mu.Lock()
	defer tm.mu.Unlock()
	t.Duration = time.Since(start).Seconds()
	if err != nil {
		t.Status = StatusFailed
		t.Error = err.Error()
	} else {
		t.Status = StatusFinished
		t.Result = result
	}
}

func (tm *TaskManager) GetTask(id string) (*Task, bool) {
	tm.mu.RLock()
	defer tm.mu.RUnlock()
	task, ok := tm.tasks[id]
	return task, ok
}

func (tm *TaskManager) GetTasks() map[string]*Task {
	tm.mu.RLock()
	defer tm.mu.RUnlock()
	tasks := tm.tasks
	return tasks
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
