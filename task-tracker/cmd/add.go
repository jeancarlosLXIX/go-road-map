package cmd

import (
	"encoding/json"
	"os"
	"task-tracker/internal/models"
)

// Recibes a task and a file name to add into the JSON.
func AddTask(task models.Task, fileName string) {
	var tasks []models.Task

	// Read file
	file, err := os.Open(fileName)
	if err == nil {
		json.NewDecoder(file).Decode(&tasks)
		file.Close()
	}

	// Add new task
	tasks = append(tasks, task)

	// Rewrite file
	file, _ = os.Create(fileName)
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	encoder.Encode(tasks)
}
