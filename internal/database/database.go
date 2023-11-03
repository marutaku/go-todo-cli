package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/marutaku/todo/internal/todo"
	_ "github.com/mattn/go-sqlite3"
)

func Connect(sqliteFilePath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", sqliteFilePath)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	if err := createTaskTable(db); err != nil {
		panic(err)
	}
	return db, nil
}

func createTaskTable(db *sql.DB) error {
	statement := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		status TEXT CHECK( status IN ('Todo','InProgress','Done') ) NOT NULL DEFAULT 'Todo',
		deadline TIMESTAMP DEFAULT NULL,
		createdAt TIMESTAMP NOT NULL DEFAULT (DATETIME('now', 'localtime'))
	)`
	_, err := db.Exec(statement)
	return err
}

func InsertTask(db *sql.DB, title string, deadline time.Time, withDeadline bool) (*todo.Task, error) {

	var statement string
	if withDeadline {
		statement = fmt.Sprintf(`INSERT INTO tasks(title, deadline) VALUES (
			"%s",
			"%s"
		)`, title, deadline)
	} else {
		statement = fmt.Sprintf(`INSERT INTO tasks(title) VALUES (
			"%s"
		)`, title)
	}

	result, err := db.Exec(statement)
	if err != nil {
		return nil, err
	}
	taskId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	task := &todo.Task{
		Id:        taskId,
		Title:     title,
		Status:    todo.Todo,
		CreatedAt: time.Now(),
	}
	if withDeadline {
		task.Deadline = deadline
	}
	return task, nil
}

func Close(db *sql.DB) error {
	return db.Close()
}
