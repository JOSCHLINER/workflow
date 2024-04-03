package handler

import (
	"errors"
	"fmt"
	"strconv"
	"workflow/internal/database"
	"workflow/internal/tableformatter"
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

	printTasks(tasks)
}

func ListActiveTasks() {

	tasks, err := dB.GetActiveTasks()
	if err != nil {
		fmt.Println("Error retrieving tasks from database!", err)
	}

	printTasks(tasks)
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

func printTasks(tasks []database.Task) {
	formatter := tableformatter.NewTableFormatter()

	taskList := tasksIntoListWithDesc(tasks)
	fmt.Println(formatter.ConstructTable(taskList))
}

func tasksIntoListWithDesc(tasks []database.Task) [][]string {

	taskList := [][]string{{"Id:", "Description:", "Time Spent:"}}
	for _, task := range tasks {
		items := []string{strconv.FormatInt(task.ID, 10), task.Name, strconv.Itoa(int(task.TimeSpent))}
		taskList = append(taskList, items)
	}

	return taskList
}
