package todo

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Status = string

const (
	Todo       = Status("Todo")
	InProgress = Status("InProgress")
	Done       = Status("Done")
)

type Task struct {
	Id        int64
	Title     string
	Status    Status
	Deadline  sql.NullTime
	CreatedAt time.Time
}

type TaskInput struct {
	Title          string
	DeadlineString string
	WithDeadline   bool
}

func (t Task) ToDescriptionString() string {
	taskDescription := fmt.Sprintf("title: %s, status: %s", t.Title, t.Status)
	if t.Deadline.Valid {
		fmt.Println(t.Deadline)
		taskDescription += fmt.Sprintf(", deadline: " + t.Deadline.Time.Format("2006-01-02"))
	}
	return taskDescription

}

func (t *Task) ChangeStatus(status Status) {
	t.Status = status
}

func Create(db *sql.DB, input TaskInput) (*Task, error) {
	if input.Title == "" {
		return nil, errors.New("title is required")
	}
	var deadline time.Time
	var err error
	statement := "INSERT INTO tasks(title) VALUES ($1)"
	if input.WithDeadline {
		deadline, err = time.Parse("2006-01-02", input.DeadlineString)
		if err != nil {
			panic(err)
		}
		statement = "INSERT INTO tasks(title, deadline) VALUES ($1,$2)"
	}
	result, err := db.Exec(statement, input.Title, deadline)
	if err != nil {
		return nil, err
	}
	taskId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	task := &Task{
		Id:        taskId,
		Title:     input.Title,
		Status:    Todo,
		CreatedAt: time.Now(),
	}
	if input.WithDeadline {
		task.Deadline = sql.NullTime{
			Time:  deadline,
			Valid: true,
		}
	}
	return task, nil
}

func ListTasks(db *sql.DB, showAllTasks bool) ([]Task, error) {
	var statement string
	if showAllTasks {
		statement = `SELECT * FROM tasks`
	} else {
		statement = `SELECT * FROM tasks WHERE status != 'Done'`
	}
	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	var tasks []Task
	for rows.Next() {
		var task Task
		if err = rows.Scan(&task.Id, &task.Title, &task.Status, &task.Deadline, &task.CreatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
