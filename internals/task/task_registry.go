package task

type TaskRegistry struct {
	tasks map[string]Task
}

func NewTaskRegistry() *TaskRegistry {
	return &TaskRegistry{
		tasks: make(map[string]Task),
	}
}

func (tr *TaskRegistry) RegisterTask(name string, task Task) {
	tr.tasks[name] = task
}
