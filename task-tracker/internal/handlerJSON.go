package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"task-tracker/internal/models"
	"time"
)

// Re-write the whole file and save it
func SaveTasks(tasks []models.Task, fileName string) {
	file, _ := os.Create(fileName)
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	encoder.Encode(tasks)
}

// This will check for a json files and if it doens't exist it will be created
func DoesExists(fileName string) {

	_, err := os.Stat(fileName)

	if err == nil {
		return
	}

	file, _ := os.Create(fileName)
	defer file.Close()

}

// get all the data from files
func GetTasks(fileName string) ([]models.Task, error) {
	var tasks []models.Task

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// assign ID to the task
func AssigID(tasks []models.Task) uint16 {
	var newID uint16
	taken := make(map[uint16]bool)

	// gettting taken IDs
	for _, task := range tasks {
		taken[task.ID] = true
	}

	// verifying
	for i := range uint16(65535) {
		// value, true or false
		_, ok := taken[i]

		if ok == false {
			newID = i
			break
		}
	}

	return newID
}

func PrintData(tasks []models.Task) {

	for _, t := range tasks {
		print := `
ID: %d
Description: %s
Status: %s
Created at: %s
Updated at: %s
***************
`
		fmt.Printf(print,
			t.ID,
			t.Description,
			t.Status,
			t.CreatedAt.Format("02 Jan 2006 15:04"),
			t.UpdatedAt.Format("02 Jan 2006 15:04"))

		// Format("2006-01-02"))       // 2026-03-31
		// Format("02/01/2006"))       // 31/03/2026
		// Format("02 Jan 2006"))      // 31 Mar 2026
		// Format("15:04"))            // 14:30
		// Format("2006-01-02 15:04")) // full datetime
	}
}

// Filter all the stasts by status
func FilterTask(tasks []models.Task, status string) []models.Task {
	var filterdTask []models.Task

	for _, task := range tasks {
		if task.Status == status {
			filterdTask = append(filterdTask, task)
		}
	}

	return filterdTask
}

func UpdateAt(ID uint16, data []models.Task) ([]models.Task, error) {

	for i, task := range data {
		if task.ID == ID {
			data[i].UpdatedAt = time.Now()
			return data, nil
		}
	}

	return nil, fmt.Errorf("task not found")
}

func Mark(ID uint16, data []models.Task, mark string) ([]models.Task, error) {

	for i, task := range data {
		if task.ID == ID {
			data[i].UpdatedAt = time.Now()
			data[i].Status = mark
			return data, nil
		}
	}

	return nil, fmt.Errorf("task not found")
}

// Delete task from our file
func DeleteTask(tasks []models.Task, id uint16) ([]models.Task, error) {
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return tasks, nil
		}
	}
	return tasks, fmt.Errorf("task not found")
}
