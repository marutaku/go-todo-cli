/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/marutaku/todo/internal/database"
	"github.com/marutaku/todo/internal/todo"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete task",
	Long:  `Delete task`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			ExitWithError(errors.New("task id undefined"))
		}
		db, err := database.Connect("./task.db")
		if err != nil {
			ExitWithError(errors.New("database not found. make sure your database file is right path"))
		}
		defer db.Close()
		taskId, err := strconv.Atoi(args[0])
		if err != nil {
			ExitWithError(fmt.Errorf("%s is not a valid task id", args[0]))
		}
		err = todo.Delete(db, taskId)
		if err != nil {
			ExitWithError(err)
		}
		fmt.Printf("Delete task id: %d\n", taskId)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
