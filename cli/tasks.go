package cli

import (
	"fmt"
	"strings"
)

type Task struct {
	Id          uint
	Description string
}

type UserTaskList struct {
	Counter uint    `json:"counter"`
	List    []*Task `json:"list"`
}

func (usertl *UserTaskList) CreateNewTask(description string) string {
	usertl.Counter++

	newTask := &Task{
		Id:          usertl.Counter,
		Description: description,
	}

	usertl.List = append(usertl.List, newTask)

	UpdateFile(usertl)

	return fmt.Sprintf("Task added successfully (ID: %d)", newTask.Id)
}

func (usertl *UserTaskList) GetAllTasks() string {
	var sb strings.Builder
	sb.WriteString("Listing all tasks:\n")

	for _, task := range usertl.List {
		sb.WriteString(
			fmt.Sprintf("%d - %s\n", task.Id, task.Description),
		)
	}

	return sb.String()
}

func (usertl *UserTaskList) GetTaskById(id uint) *Task {

	for _, task := range usertl.List {
		if task.Id == id {
			return task
		}
	}

	return nil
}

func createNewUserTaskList() *UserTaskList {
	return &UserTaskList{
		0,
		[]*Task{},
	}
}
