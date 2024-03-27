package handler

import (
	"errors"
	"fmt"
	"os"
	"text/tabwriter"
	"workflow/internal/database"
)

const path string = "./data.sqlite"

var dB *database.Database

func init() {
	db, err := database.Connect(path)
	if err != nil {
		panic("Saved file can not be accessed, quitting...")
	}

	dB = db
}

func Initialize() {
	if err := dB.CreateTasksTable(); err != nil {
		fmt.Println(errors.New("Could not initialize the program!"), err)
		return
	}

	fmt.Println("Program successfully initialized.")
}

func ListAllTasks() {

	tasks, err := dB.GetAllTasks()
	if err != nil {
		fmt.Println("Error retrieving tasks from database!", err)
	}

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', tabwriter.Debug)
	defer w.Flush()

	fmt.Fprintln(w, "Id:\tDescription:\tTime Spent:")

	for _, task := range tasks {
		line := fmt.Sprintf("%d\t%s\t%d", task.ID, task.Name, int(task.TimeSpent))
		fmt.Fprintln(w, line)
	}
}

func ListActiveTasks() {

	tasks, err := dB.GetActiveTasks()
	if err != nil {
		fmt.Println("Error retrieving tasks from database!", err)
	}

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', tabwriter.Debug)
	defer w.Flush()

	fmt.Fprintln(w, "Id:\tDescription:\tTime Spent:")
	fmt.Fprintln(w, "%-10\t%-10\t%-10")
	for _, task := range tasks {
		line := fmt.Sprintf("%d\t%s\t%d", task.ID, task.Name, int(task.TimeSpent))
		fmt.Fprintln(w, line)
	}

	w.Flush()
}

func CreateTask(name string) {
	task := database.Task{Name: name}
	err := dB.CreateTask(&task)
	if err != nil {
		fmt.Println("Error occurred! Task couldn't be created!", err)
		return
	}

	fmt.Println("Task successfully created.")
}

func StartTask(id int) {
	task, err := getTask(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := dB.StartTask(&task); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Task has been started.")
}

func StopTask(id int) {
	task, err := getTask(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := dB.StopTask(&task); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Task has been stopped.")
}

func getTask(id int) (database.Task, error) {
	task := database.Task{ID: int64(id)}

	err := dB.GetTask(&task)
	if err != nil {
		return task, errors.New("Task not found!")
	}

	return task, nil
}
