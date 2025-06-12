package task

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
