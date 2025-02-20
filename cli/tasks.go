package cli

import "fmt"

type Task struct {
	id          uint
	description string
}

type UserTaskList struct {
	counter uint
	list    []*Task
}

func (usertl *UserTaskList) CreateNewTask(description string) string {
	usertl.counter++
	newTask := &Task{
		id:          usertl.counter,
		description: description,
	}
	usertl.list = append(usertl.list, newTask)

	return fmt.Sprintf("Task added successfully (ID: %d)", newTask.id)
}

func CreateNewUserTaskList() *UserTaskList {
	return &UserTaskList{
		0,
		[]*Task{},
	}
}
