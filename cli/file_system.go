package cli

import (
	"encoding/json"
	"log"
	"os"
)

const (
	filePath = "./data/task-list.json"
)

func isFileExists() bool {
	_, err := os.Stat(filePath)

	if err != nil {
		return false
	}

	return true
}

func CreateFile() error {
	if isFileExists() {
		return nil
	}

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		return err
	}

	UpdateFile(createNewUserTaskList())

	file.Close()
	return nil
}

func UpdateFile(userTaskList *UserTaskList) {

	jsonData, err := json.Marshal(userTaskList)

	if err != nil {
		log.Fatal("can't update file")
	}

	os.WriteFile(filePath, jsonData, 0664)
}

func ReadFile() *UserTaskList {
	file, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal("can't read file")
	}

	data := &UserTaskList{}

	json.Unmarshal(file, data)

	return data
}
