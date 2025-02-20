package cli

import (
	"os"
)

func Reader(userTaskList *UserTaskList) string {
	args := os.Args[1:]

	switch args[0] {
	case "add":
		answer := userTaskList.CreateNewTask(args[1])
		return answer
	default:
		return "command no valid"
	}
}
