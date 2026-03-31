package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"task-tracker/cmd"
	"task-tracker/internal"
	"task-tracker/internal/models"
	"time"
)

func main() {
	const FILE_NAME string = "tasks.json"
	const (
		IN_PROGRESS string = "⏳"
		DONE        string = "✅"
	)
	var task models.Task
	internal.DoesExists(FILE_NAME)

	// using standard package for the commands
	commands := os.Args[1:]

	if len(commands) == 0 {
		fmt.Println("No command provided")
		return
	}

	main_command := strings.ToLower(commands[0])

	data, err := internal.GetTasks(FILE_NAME)

	// if the list is empty we can't print nor update
	condition := (err == io.EOF && main_command == "list") ||
		(err == io.EOF && main_command == "update")

	if condition { // End Of File so it means our file is empty
		fmt.Printf("Empty %s.\n", FILE_NAME)
		return
	}

	switch main_command {
	case "add":
		commands_size := int(len(commands))
		if commands_size < 2 {
			fmt.Println("add command needs a task")
			fmt.Println("Example: add 'Buy groceries.'")
			return
		}

		if commands_size > 3 {
			fullString := strings.Join(commands[1:], " ")
			task.Description = fullString
		} else {
			task.Description = commands[1]
		}

		task.ID = internal.AssigID(data)
		task.Status = IN_PROGRESS
		task.CreatedAt = time.Now()
		task.UpdatedAt = time.Now()
		cmd.AddTask(task, FILE_NAME)

	case "list":
		internal.PrintData(data)

	case "update":
		n, err := strconv.Atoi(commands[1])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		data, err := internal.UpdateAt(uint16(n), data)

		if err != nil {
			fmt.Println("No task with that ID " + commands[1])
			return
		}

		reader := bufio.NewReader(os.Stdin)

		fmt.Print("New Description: ")
		newDesc, _ := reader.ReadString('\n')
		newDesc = strings.TrimSpace(newDesc)

		for i, t := range data {
			if t.ID == uint16(n) {
				data[i].Description = newDesc
				data[i].UpdatedAt = time.Now()
				break
			}
		}

		internal.SaveTasks(data, FILE_NAME)

	case "delete":
		n, err := strconv.Atoi(commands[1])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		data, err = internal.DeleteTask(data, uint16(n))

		if err != nil {
			fmt.Print("Check your ID")
			return
		}
		fmt.Println("Task deleted " + commands[1])
		internal.SaveTasks(data, FILE_NAME)

	case "done-task":
		tasks := internal.FilterTask(data, DONE)
		internal.PrintData(tasks)

	case "undone-task":
		tasks := internal.FilterTask(data, IN_PROGRESS)
		internal.PrintData(tasks)

	case "mark-done":
		n, err := strconv.Atoi(commands[1])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		data, err := internal.Mark(uint16(n), data, DONE)
		if err != nil {
			fmt.Println("No task with that ID " + commands[1])
			return
		}

		internal.SaveTasks(data, FILE_NAME)

	case "mark-undone":
		n, err := strconv.Atoi(commands[1])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		data, err := internal.Mark(uint16(n), data, IN_PROGRESS)
		if err != nil {
			fmt.Println("No task with that ID " + commands[1])
			return
		}

		internal.SaveTasks(data, FILE_NAME)

	case "list-done":
		task := internal.FilterTask(data, DONE)

		internal.PrintData(task)

	case "list-undone":
		task := internal.FilterTask(data, IN_PROGRESS)

		internal.PrintData(task)

	default:
		fmt.Println("Nothing to do here.")

	}

}
