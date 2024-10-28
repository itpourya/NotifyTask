package tasks

import (
	"encoding/json"
	"log"
)

func (t *Task) DeleteTask(taskId int) bool {
	_ = loadJsonData()

	id := getItem(taskId)
	newTask := removeIndex(Tasks, id)

	jsonData, err := json.Marshal(newTask)
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
