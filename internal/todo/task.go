package todo

import (
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

func (t Task) ToDescriptionString() string {
	return fmt.Sprintf("title: " + t.Title + ", status: " + t.Status + ", deadline: " + t.Deadline.Format("2006-01-02 15:04:05"))
}

func (t *Task) ChangeStatus(status Status) {
	t.Status = status
}
