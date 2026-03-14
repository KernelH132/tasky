package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	usage := `
	---- Usage ----
	- add [task]: Create a new task
	- ls         : List all tasks
	- usage      : Show this menu
	- exit       : Exit the CLI
	- rm taskID  : Remeove a task
	- write taskID newTask   : Edit a task
	- clear      : Delete all tasks
	- save       : Save tasks to a file
	- check      : Mark as complete
	- ls -c      : lists all checked tasks // good feature but not necessary
	- ls -u      : lists all uncompleted // good feature but not necessary
	`

	tasks := make(map[int]string)
	count := 1
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to Tasky!")
	fmt.Println(usage)

	for {
		fmt.Print("tasky> ")
		if !scanner.Scan() {
			break
		}

		inputs := scanner.Text()
		parts := strings.Fields(inputs)

		if len(parts) == 0 {
			continue
		}

		command := parts[0]

		switch command {
		case "add":
			if len(parts) < 2 {
				fmt.Println("Error: Please provide a task description.")
				continue
			}

			taskDesc := strings.Join(parts[1:], " ")
			tasks[count] = taskDesc
			fmt.Printf("Added task #%d\n", count)
			count++
		case "ls":
			if len(tasks) == 0 {
				fmt.Println("No tasks yet!")
			}
			for i, task := range tasks {
				fmt.Printf("%d: %s\n", i, task)
			}
		case "rm":
			if len(tasks) == 0 {
				fmt.Println("No tasks to delete!")
			}
			if len(parts) > 2 {
				fmt.Println("Error: Can only have one argument.")
			}
			id, _ := strconv.Atoi(parts[1])
			delete(tasks, id)
			fmt.Println("Task removed!")
		case "check":
			if len(parts) > 2 {
				fmt.Println("Error: Can only have one argument.")
			}

			if len(parts) < 2 {
				fmt.Println("Error: Please provide a task ID.")
				continue
			}

			id, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Error: ID must be a number.")
			}

			if task, exists := tasks[id]; exists {
				tasks[id] = task + " (completed)"
				fmt.Println("Task marked as completed!")
			} else {
				fmt.Println("Task doesn't exist.")
			}

		case "write":
			id, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Error: ID must be a number.")
				continue
			}

			if len(parts) < 3 {
				fmt.Println("Error: Usage: write [id] [new description].")
				continue
			}

			if id > len(tasks) {
				fmt.Println("Task dosen't exist.")
			}

			taskDesc := strings.Join(parts[2:], " ")

			// if id > len(tasks) {
			// 	fmt.Println("Task dosen't exist.")
			// }

			// for i := range tasks {
			// 	if i == id {
			// 		tasks[id] = taskDesc
			// 	}

			// }
			if _, exists := tasks[id]; exists {
				tasks[id] = taskDesc
				fmt.Println("Task edited!")
			} else {
				fmt.Println("Task ID not found.")
			}

		case "clear":
			clear(tasks)
			fmt.Printf("Deleted all tasks.")

		case "save":
			file, err := os.Create("tasks.json")
			if err != nil {
				fmt.Println(err)
			}

			defer file.Close()

			encoder := json.NewEncoder((file))
			encoder.Encode(tasks)

		case "exit":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Command not identified! Enter 'usage' to see options.")
		}

	}

}
