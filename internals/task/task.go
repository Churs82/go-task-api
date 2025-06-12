package task

import (
	"fmt"
)

var taskRegistry *TaskRegistry

func init() {
	taskRegistry = NewTaskRegistry()
	taskRegistry.RegisterTask("io", &IoTask{})
}

type Task interface {
	Run() (string, error)
}

func RunTask(name string) (string, error) {
	task, ok := taskRegistry.tasks[name]
	if !ok {
		return "", fmt.Errorf("task not found: %s", name)
	}
	return task.Run()
}
