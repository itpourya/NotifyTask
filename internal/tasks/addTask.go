package tasks

import (
	"encoding/json"
	"log"
	"time"
)

func (t *Task) AddTask(tasks string) bool {
	_ = loadJsonData()
	var id int = 0
	var task = &Task{}

	for _, item := range Tasks {
		id = item.ID
	}

	if id == 0 {
		task.ID = 1
		task.Name = tasks
		task.Status = false
		task.CreatedAt = time.Now()
	} else {
		task.ID = id + 1
		task.Name = tasks
		task.Status = false
		task.CreatedAt = time.Now()
	}

	Tasks = append(Tasks, *task)

	jsonData, err := json.Marshal(Tasks)
	if err != nil {
		log.Println("Faild to marshal task data.")
		return false
	}

	status := writeJsonData(jsonData)
	if !status {
		return false
	}

	return true
}
