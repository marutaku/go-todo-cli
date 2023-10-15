package main

import (
	"fmt"
	"time"

	"github.com/marutaku/todo/internal/database"
)

func main() {
	db, err := database.Connect("./task.db")
	if err != nil {
		panic(err)
	}
	defer database.Close(db)

	if err := database.CreateTaskTable(db); err != nil {
		panic(err)
	}

	deadline := time.Date(2023, 10, 15, 12, 0, 0, 0, time.Local)
	task, err := database.InsertTask(db, "洗濯物を干す", deadline)
	if err != nil {
		panic(err)
	}

	fmt.Println(task.ToDescriptionString())
}
