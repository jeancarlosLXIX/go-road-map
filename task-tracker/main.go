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
		messages("Task added.")

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
		messages("Task has been updated.")

	case "delete":
		if lessThanExpected(len(commands), 2) {
			fmt.Println("Need to add an id")
			return
		}
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
		internal.SaveTasks(data, FILE_NAME)
		messages("Task deleted " + commands[1])
	case "done-task":
		tasks := internal.FilterTask(data, DONE)
		internal.PrintData(tasks)
		messages("Task mark as done.")

	// case "undone-task":
	// 	tasks := internal.FilterTask(data, IN_PROGRESS)
	// 	internal.PrintData(tasks)
	// 	messages("Task mark as undone.")

	case "mark-done":
		if lessThanExpected(len(commands), 2) {
			fmt.Println("Need to add the ID of the task that has been done.")
			return
		}
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
		messages("Tasked mark as Done")

	case "mark-undone":
		if lessThanExpected(len(commands), 2) {
			fmt.Println("Need to add the ID of the unmark task")
			return
		}
		n, err := strconv.Atoi(commands[1])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		data, err := internal.Mark(uint16(n), data, IN_PROGRESS)
		if err != nil {
			fmt.Println("No task with ID " + commands[1])
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
		printHelp()

	}

}

func printHelp() {
	print := `
	add <Your task here>
	delete <ID>
	update <ID>
	list prints your
	list-undone print your undone tasks
	list-done prints your done tasks
	`
	fmt.Println("Commands:\n" + print)
}

// In case the arguments for the commands are less than expected
func lessThanExpected(n int, atLeast int) bool {
	if n < atLeast {
		return true
	}
	return false
}

// Show messages after an action
func messages(msg string) {
	fmt.Println(msg)
}
