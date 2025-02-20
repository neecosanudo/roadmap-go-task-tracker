package cli

import "fmt"

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

func createNewUserTaskList() *UserTaskList {
	return &UserTaskList{
		0,
		[]*Task{},
	}
}
