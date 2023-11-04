package todo

import (
	"database/sql"
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
	Deadline  time.Time
	CreatedAt time.Time
}

type TaskInput struct {
	Title          string
	DeadlineString string
	WithDeadline   bool
}

func (t Task) ToDescriptionString() string {
	return fmt.Sprintf("title: " + t.Title + ", status: " + t.Status + ", deadline: " + t.Deadline.Format("2006-01-02"))
}

func (t *Task) ChangeStatus(status Status) {
	t.Status = status
}

func Create(input TaskInput, db *sql.DB) (*Task, error) {
	var deadline time.Time
	var statement string
	if input.WithDeadline {
		deadline, err := time.Parse("2006-01-02", input.DeadlineString)
		if err != nil {
			panic(err)
		}
		statement = fmt.Sprintf(`INSERT INTO tasks(title, deadline) VALUES (
			"%s",
			"%s"
		)`, input.Title, deadline)
	} else {
		statement = fmt.Sprintf(`INSERT INTO tasks(title) VALUES (
			"%s"
		)`, input.Title)
	}
	result, err := db.Exec(statement)
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
		task.Deadline = deadline
	}
	return task, nil
}
