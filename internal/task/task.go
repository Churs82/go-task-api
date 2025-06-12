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
	ID        string     `json:"id"`
	Status    TaskStatus `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	Duration  float64    `json:"duration"`
	Result    string     `json:"result,omitempty"`
	Error     string     `json:"error,omitempty"`
	TaskType  string     `json:"task_type,omitempty"`
}

func (t *Task) Run() (string, error) {
	return registry.taskTypes[t.TaskType].Run()
}

type TaskRegistry struct {
	taskTypes map[string]Task_i
}

func (tr *TaskRegistry) RegisterTask(name string, i Task_i) {
	tr.taskTypes[name] = i
}

func NewTaskRegistry() *TaskRegistry {
	return &TaskRegistry{
		taskTypes: make(map[string]Task_i),
	}
}

var registry = NewTaskRegistry()

func init() {
	// Register tasks
	// IO Task
	registry.RegisterTask("io", &IoTask{})
	// To add more task types register them here
	// .....
}
