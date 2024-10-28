package tasks

import (
	"encoding/json"
	"log"
	"time"
)

func (t *Task) CompleteTask(taskId int) bool {
	_ = loadJsonData()

	id := getItem(taskId)

	Tasks[id].CompletedAt = time.Now()
	Tasks[id].Status = true

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
