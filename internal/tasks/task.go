package tasks

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type TodoUtils interface {
	AddTask(task string) bool
	ShowTask() []string
	CompleteTask(taskId int) bool
	DeleteTask(taskId int) bool
}

type Task struct {
	ID          int
	Name        string
	Status      bool
	CreatedAt   time.Time
	CompletedAt time.Time
	NotifAt     time.Time
}

var Tasks []Task

var JsonFile = ".todo.json"

func loadJsonData() bool {
	file, err := os.ReadFile(JsonFile)
	if err != nil {
		log.Println("Faild to read todo.json file.")
		log.Println("Create todo.json file.")
	}

	err = json.Unmarshal(file, &Tasks)
	if err != nil {
		log.Println("Faild to unmarshal json file.")
		return false
	}

	return true
}

func writeJsonData(jsonData []byte) bool {
	err := os.WriteFile(JsonFile, jsonData, 0666)
	if err != nil {
		log.Println("Faild to marshal task data.")
		return false
	}

	return true
}

func removeIndex(s []Task, index int) []Task {
	return append(s[:index], s[index+1:]...)
}

func getItem(taskID int) int {
	var id = 0
	for _, item := range Tasks {
		id += 1
		if item.ID == taskID {
			return id - 1
		}
	}

	return 0
}
