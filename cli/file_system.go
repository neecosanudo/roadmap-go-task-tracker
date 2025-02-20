package cli

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	dirPath  = "./data"
	filePath = dirPath + "/task-list.json"
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

	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	err = updateFile(createNewUserTaskList())
	if err != nil {
		return fmt.Errorf("failed to update file: %w", err)
	}

	file.Close()
	return nil
}

func updateFile(userTaskList *UserTaskList) error {
	jsonData, err := json.Marshal(userTaskList)

	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	return os.WriteFile(filePath, jsonData, 0664)
}

func ReadFile() (*UserTaskList, error) {
	file, err := os.ReadFile(filePath)

	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	data := &UserTaskList{}

	err = json.Unmarshal(file, data)

	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal data: %w", err)
	}

	return data, nil
}
