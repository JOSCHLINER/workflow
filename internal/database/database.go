package database

import (
	"database/sql"
	"errors"
	"time"
	"fmt"	
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db *sql.DB
}

type Task struct {
	ID        int64
	Name      string
	TimeSpent time.Time
	Active    bool
	StartTime time.Time
}

func Connect(path string) (*Database, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return &Database{db}, nil
}

func (d *Database) CreateTasksTable() error {
	sql := `CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title VARCHAR(255) NOT NULL,
		time_spent DATETIME,
		active BOOLEAN DEFAULT 0,
		start_time DATETIME
		);`

	_, err := d.db.Exec(sql)
	return err
}

func (d *Database) CreateTask(task *Task) error {
	res, err := d.db.Exec(`INSERT INTO tasks(title) VALUES(?)`, task.Name)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	task.ID = id

	return nil
}

func (d *Database) GetTask(task *Task) error {

	res := d.db.QueryRow(`SELECT * FROM tasks WHERE id = ?`, task.ID)

	if err := res.Scan(&task.ID, &task.Name, &task.TimeSpent, &task.Active, &task.StartTime); err != nil {
		return errors.New("No task with such id!")
	}

	return nil
}

func (d *Database) GetAllTasks() ([]Task, error) {

	res, err := d.db.Query(`SELECT * FROM tasks`)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	var tasks []Task
	sqlStartTime, sqlTimeSpent := "", ""
	for res.Next() {
		var task Task
		if err := res.Scan(&task.ID, &task.Name, &sqlStartTime, &task.Active, &sqlTimeSpent); err != nil {
			return nil, err
		}
		
		task.TimeSpent = convertToTime(sqlTimeSpent)
		task.StartTime = convertToTime(sqlStartTime)

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (d *Database) StartTask(task *Task) error {

	res, err := d.db.Exec(`UPDATE tasks SET active = 1,
				start_time = CURRENT_TIMESTAMP
				WHERE id = ?`, task.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if rowsAffected == 0 {
		return errors.New("Invalid task Id!")
	}

	task.Active = true
	return nil
}

func (d *Database) StopTask(task *Task) error {

	res, err := d.db.Exec(`UPDATE tasks 
	SET active = 0,
	    time_spent = time(
		    strftime('%s', time_spent) + 
		    (strftime('%s', 'now') - strftime('%s', start_time)), 'unixepoch' )
	WHERE active = 1 and id = ?;`, task.ID)

	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if rowsAffected == 0 {
		return errors.New("No such task or task is not active")
	}

	return nil
}


func convertToTime(datetime string) time.Time {
	layout := "2006-01-02T15:04:05Z"

	parsedTime, err := time.Parse(layout, datetime)
	if err != nil {
		fmt.Println(datetime)
		return time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
	}


	fmt.Println(parsedTime)
	return parsedTime
}


