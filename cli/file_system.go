package cli

import (
	"encoding/json"
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

	_, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(CreateNewUserTaskList())

	if err != nil {
		return err
	}

	os.WriteFile(filePath, jsonData, 0664)

	return nil
}
