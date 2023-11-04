/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/marutaku/todo/internal/database"
	"github.com/marutaku/todo/internal/todo"
	"github.com/spf13/cobra"
)

type AddCommandOption struct {
	deadline string
}

var option = AddCommandOption{}

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
			panic(err)
		}
		defer database.Close(db)
		withDeadline := option.deadline != ""
		input := todo.TaskInput{Title: args[0], DeadlineString: option.deadline, WithDeadline: withDeadline}
		task, err := todo.Create(input, db)
		if err != nil {
			panic(err)
		}
		fmt.Println(task.ToDescriptionString())
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&option.deadline, "deadline", "d", "", "Task deadline")
}
