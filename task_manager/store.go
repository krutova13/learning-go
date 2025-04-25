package task_manager

type Store[T any] interface {
	LoadData(item T)
	GetData() []T
}

type TaskStore struct {
	tasks []*Task
}

func NewTaskStore() *TaskStore {
	return &TaskStore{
		tasks: make([]*Task, 0),
	}
}

func (t *TaskStore) LoadData(task *Task) {
	t.tasks = append(t.tasks, task)
}

func (t *TaskStore) GetData() []*Task {
	return t.tasks
}
