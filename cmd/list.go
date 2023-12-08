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

type ListTasksOption struct {
	showAll bool
}

var listTaskOption = ListTasksOption{}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks",
	Long: `List the tasks that haven't been completed. 
By adding the '--all' option, you can view all tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := database.Connect("./task.db")
		if err != nil {
			ExitWithError(errors.New("database not found. make sure your database file is right path"))
		}
		defer database.Close(db)
		tasks, err := todo.ListTasks(db, listTaskOption.showAll)
		if err != nil {
			panic(err)
		}
		for _, task := range tasks {
			fmt.Println(task.ToDescriptionString())
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&listTaskOption.showAll, "all", "a", false, "Show also finished tasks")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
