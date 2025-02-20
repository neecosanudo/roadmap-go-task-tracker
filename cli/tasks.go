package cli

import (
	"fmt"
	"strings"
)

const (
	STATUS_TO_DO = iota
	STATUS_IN_PROGRESS
	STATUS_DONE
)

type Task struct {
	Id          uint
	Description string
	Status      uint
}

type UserTaskList struct {
	Counter uint    `json:"counter"`
	List    []*Task `json:"list"`
}

func (usertl *UserTaskList) createNewTask(description string) string {
	usertl.Counter++

	newTask := &Task{
		Id:          usertl.Counter,
		Description: description,
		Status:      STATUS_TO_DO,
	}

	usertl.List = append(usertl.List, newTask)

	updateFile(usertl)

	return fmt.Sprintf("Task added successfully (ID: %d)", newTask.Id)
}

func (usertl *UserTaskList) getAllTasks() string {
	var sb strings.Builder
	sb.WriteString("Listing all tasks:\n\n")

	for _, task := range usertl.List {
		sb.WriteString(
			fmt.Sprintf("Task %d - %s - Status: %s\n", task.Id, task.Description, assertStatusByNumber(task.Status)),
		)
	}

	return sb.String()
}

func (usertl *UserTaskList) getTaskById(id uint) (*Task, int) {

	for i, task := range usertl.List {
		if task.Id == id {
			return task, i
		}
	}

	return nil, 0
}

func (usertl *UserTaskList) deleteTaskById(id uint) string {
	task, index := usertl.getTaskById(id)

	if task == nil {
		return "task doesn't exist"
	}

	usertl.List = assertRemoveTask(usertl.List, index)

	updateFile(usertl)

	return fmt.Sprintf("Task deleted successfully (ID: %d)", task.Id)
}

func (usertl *UserTaskList) updateDescriptionTaskById(id uint, newDescription string) string {
	task, index := usertl.getTaskById(id)

	if task == nil {
		return fmt.Sprintf("Task with ID %d doesn't exist", id)
	}

	usertl.List[index].Description = newDescription

	updateFile(usertl)

	return fmt.Sprintf("Task updated successfully (ID: %d)", task.Id)
}

func (usertl *UserTaskList) getAllTasksByStatus(statusId uint) string {
	var sb strings.Builder
	title := fmt.Sprintf("All tasks with status:%s \n", assertStatusByNumber(statusId))
	sb.WriteString(title)

	for _, task := range usertl.List {
		if task.Status == statusId {
			sb.WriteString(
				fmt.Sprintf("Task %d - Description: %s\n", task.Id, task.Description),
			)
		}
	}

	return sb.String()
}

func (usertl *UserTaskList) updateStatusTaskInto(id uint, status uint) string {
	task, index := usertl.getTaskById(id)

	usertl.List[index].Status = status

	updateFile(usertl)

	return fmt.Sprintf("Task status updated successfully (ID: %d)", task.Id)
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

func assertStatusByNumber(number uint) string {
	switch number {
	case 1:
		return "In Progress"
	case 2:
		return "Done"
	default:
		return "To Do"
	}
}
