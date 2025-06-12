package task

import (
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type TaskStatusJson struct {
	Status    TaskStatus `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	Duration  float64    `json:"duration,omitempty"`
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

func (tm *TaskManager) CreateTask(taskType string) map[string]*Task {
	id := strconv.FormatInt(time.Now().UnixNano()+rand.Int63n(1000), 36)
	task := &Task{
		Status:    StatusPending,
		CreatedAt: time.Now(),
		TaskType:  taskType,
	}
	tm.mu.Lock()
	tm.tasks[id] = task
	tm.mu.Unlock()
	res := make(map[string]*Task, 1)
	res[id] = task
	go tm.runTask(task)
	return res
}

func (tm *TaskManager) runTask(t *Task) {
	tm.mu.Lock()
	t.Status = StatusRunning
	t.StartedAt = time.Now()
	tm.mu.Unlock()

	// Run in parallel without blocking
	result, err := t.Run()

	tm.mu.Lock()
	defer tm.mu.Unlock()
	t.Duration = time.Since(t.StartedAt).Seconds()
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

func (tm *TaskManager) GetTaskStatus(id string) (TaskStatusJson, bool) {
	tm.mu.RLock()
	defer tm.mu.RUnlock()
	s := TaskStatusJson{}
	t, ok := tm.tasks[id]
	if ok {
		s.Status = t.Status
		s.CreatedAt = t.CreatedAt
		switch t.Status {
		case StatusRunning:
			s.Duration = time.Since(t.StartedAt).Seconds()

		case StatusFinished:
			s.Duration = t.Duration

		default:
			s.Duration = time.Since(t.CreatedAt).Seconds()
		}
	}
	return s, ok
}

func (tm *TaskManager) GetTaskResult(id string) (string, bool) {
	tm.mu.RLock()
	defer tm.mu.RUnlock()
	var res string
	t, ok := tm.tasks[id]
	if ok {
		if t.Status == StatusFinished {
			res = t.Result
		} else {
			res = string(t.Status)
		}
	}
	return res, ok
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
