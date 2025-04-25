package task_manager

type TaskManager struct {
	store Store[*Task]
}

func NewTaskManager(store Store[*Task]) *TaskManager {
	return &TaskManager{store: store}
}

func (tm TaskManager) GetAllData() []*Task {
	return tm.store.GetData()
}

func (tm TaskManager) AddTask(task *Task) {
	tm.store.LoadData(task)
}
