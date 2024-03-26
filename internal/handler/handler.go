package handler

import (
	"fmt"
	"os"
	"text/tabwriter"
	"workflow/internal/database"
)

const path string = "./data/data.sqlite"

var dB *database.Database

func init() {
	db, err := database.Connect(path)
	if err != nil {
		panic("Saved file can not be accessed, quitting...")
	}

	dB = db
}

func ListAllTasks() {

	tasks, err := dB.GetAllTasks()
	if err != nil {
		fmt.Println("Error retrieving tasks from database!", err)
	}

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', tabwriter.Debug)
	for _, task := range tasks {
		line := fmt.Sprintf("%d\t%s\t%d", task.ID, task.Name, int(task.TimeSpent))
		fmt.Fprintln(w, line)
	}

	w.Flush()
}

func ListActiveTasks() {

	tasks, err := dB.GetActiveTasks()
	if err != nil {
		fmt.Println("Error retrieving tasks from database!", err)
	}

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', tabwriter.Debug)
	for _, task := range tasks {
		line := fmt.Sprintf("%d\t%s\t%d", task.ID, task.Name, int(task.TimeSpent))
		fmt.Fprintln(w, line)
	}

	w.Flush()

}
