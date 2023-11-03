/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/marutaku/todo/internal/database"
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
		var deadline time.Time
		if withDeadline {
			deadline, err = time.Parse("2006-01-02", option.deadline)
			if err != nil {
				panic(err)
			}
		}
		task, err := database.InsertTask(db, args[0], deadline, withDeadline)
		if err != nil {
			panic(err)
		}

		fmt.Println(task.ToDescriptionString())
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&option.deadline, "deadline", "d", "", "Task deadline")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
