package task

import (
	"time"
)

type TaskStatus string

const (
	StatusPending   TaskStatus = "pending"
	StatusRunning   TaskStatus = "running"
	StatusFinished  TaskStatus = "finished"
	StatusFailed    TaskStatus = "failed"
	StatusCancelled TaskStatus = "cancelled"
)

type Task_i interface {
	Run() (string, error)
}

type Task struct {
	Status    TaskStatus `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	StartedAt time.Time  `json:"started_at,omitempty"`
	Duration  float64    `json:"duration,omitempty"`
	Result    string     `json:"result,omitempty"`
	Error     string     `json:"error,omitempty"`
	TaskType  string     `json:"task_type"`
}

func (t *Task) Run() (string, error) {
	return registry.taskTypes[t.TaskType].Run()
}
