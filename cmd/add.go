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

var o = AddCommandOption{}

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

		if err := database.CreateTaskTable(db); err != nil {
			panic(err)
		}

		deadline := time.Date(2023, 10, 15, 12, 0, 0, 0, time.Local)
		task, err := database.InsertTask(db, "洗濯物を干す", deadline)
		if err != nil {
			panic(err)
		}

		fmt.Println(task.ToDescriptionString())
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&o.deadline, "deadline", "d", "", "Task deadline")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
