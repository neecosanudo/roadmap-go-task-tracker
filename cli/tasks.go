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

func (usertl *UserTaskList) GetTaskById(id uint) (*Task, int) {

	for i, task := range usertl.List {
		if task.Id == id {
			return task, i
		}
	}

	return nil, 0
}

func (usertl *UserTaskList) DeleteTaskById(id uint) string {
	task, index := usertl.GetTaskById(id)

	if task == nil {
		return "task doesn't exist"
	}

	usertl.List = assertRemoveTask(usertl.List, index)

	UpdateFile(usertl)

	return fmt.Sprintf("Task deleted successfully (ID: %d)", task.Id)
}

func (usertl *UserTaskList) UpdateTaskById(id uint, newDescription string) string {
	task, index := usertl.GetTaskById(id)

	usertl.List[index].Description = newDescription

	UpdateFile(usertl)

	return fmt.Sprintf("Task updated successfully (ID: %d)", task.Id)
}

func createNewUserTaskList() *UserTaskList {
	return &UserTaskList{
		0,
		[]*Task{},
	}
}

func assertRemoveTask(list []*Task, index int) []*Task {
	return append(list[:index], list[index+1:]...)
}
