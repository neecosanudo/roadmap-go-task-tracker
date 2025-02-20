package main

import (
	"fmt"
	"log"

	"github.com/neecosanudo/roadmap-go-task-tracker/cli"
)

func main() {
	if err := cli.CreateFile(); err != nil {
		log.Fatalf("failed to create file: %v", err)
	}

	userTaskList, err := cli.ReadFile()
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	todo := cli.Reader(userTaskList)

	fmt.Println(todo)
}
