/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/marutaku/todo/internal/database"
	"github.com/marutaku/todo/internal/todo"
	"github.com/spf13/cobra"
)

type AddCommandOption struct {
	deadline string
}

var addCommandOption = AddCommandOption{}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add your TODO task",
	Long: `Add your TODO task.
If you want to apply some deadline to the task, you can use '--deadline' option.
The deadline option should input with 'YYYY-MM-DD' format.
usage: 
	todo add "some task" --deadline 2023-01-01`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := database.Connect("./task.db")
		if err != nil {
			ExitWithError(errors.New("database not found. make sure your database file is right path"))
		}
		defer database.Close(db)
		withDeadline := addCommandOption.deadline != ""
		if len(args) == 0 {
			ExitWithError(errors.New("input task information"))
		}
		input := todo.TaskInput{Title: args[0], DeadlineString: addCommandOption.deadline, WithDeadline: withDeadline}
		task, err := todo.Create(db, input)
		if err != nil {
			panic(err)
		}
		fmt.Println(task.ToDescriptionString())
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&addCommandOption.deadline, "deadline", "d", "", "Task deadline")
}
