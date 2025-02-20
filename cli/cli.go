package cli

import (
	"fmt"
	"os"
	"strconv"
)

func Reader(userTaskList *UserTaskList) string {
	args := os.Args[1:]

	switch args[0] {
	case "add":
		if len(args) != 2 {
			return "invalid number of arguments"
		}
		return userTaskList.createNewTask(args[1])
	case "list":
		if len(args) == 1 {
			return userTaskList.getAllTasks()
		}
		if len(args) == 2 {
			status := assertNumberByStatus(args[1])
			return userTaskList.getAllTasksByStatus(status)
		}
	case "delete":
		if len(args) != 2 {
			return "invalid number of arguments"
		}
		id := assertIdAsUint(args[1])
		return userTaskList.deleteTaskById(id)
	case "update":
		if len(args) != 3 {
			return "invalid number of arguments"
		}
		id := assertIdAsUint(args[1])
		return userTaskList.updateDescriptionTaskById(id, args[2])
	case "mark-in-progress":
		if len(args) != 2 {
			return "invalid number of arguments"
		}
		id := assertIdAsUint(args[1])
		return userTaskList.updateStatusTaskInto(id, STATUS_IN_PROGRESS)
	case "mark-done":
		if len(args) != 2 {
			return "invalid number of arguments"
		}
		id := assertIdAsUint(args[1])
		return userTaskList.updateStatusTaskInto(id, STATUS_DONE)
	}

	return "command not valid"
}

func assertIdAsUint(idString string) uint {
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		fmt.Println("error:", err)
	}

	return uint(id)
}

func assertNumberByStatus(status string) uint {
	switch status {
	case "in-progress":
		return uint(1)
	case "done":
		return uint(2)
	default:
		return uint(0)
	}
}
