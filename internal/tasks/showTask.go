package tasks

func (t *Task) ShowTask() []Task {
	_ = loadJsonData()

	return Tasks
}
