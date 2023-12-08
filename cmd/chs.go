/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/marutaku/todo/internal/database"
	"github.com/marutaku/todo/internal/todo"
	"github.com/spf13/cobra"
)

// chsCmd represents the chs command
var chsCmd = &cobra.Command{
	Use:   "chs",
	Short: "Change task status",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		taskId, err := strconv.Atoi(args[0])
		if err != nil {
			ExitWithError(fmt.Errorf("%s is not a valid taskId", args[0]))
		}
		newStatus := args[1]
		if !todo.ValidateStatus(newStatus) {
			fmt.Printf("%s is not a valid status. Please input 'Todo', 'InProgress', 'Done'.\n", newStatus)
			os.Exit(1)
		}
		db, err := database.Connect("./task.db")
		if err != nil {
			panic(err)
		}
		defer database.Close(db)
		task, err := todo.FindTaskById(db, taskId)
		if err != nil {
			ExitWithError(err)
		}
		newTask, err := task.ChangeStatus(newStatus, db)
		if err != nil {
			ExitWithError(err)
		}
		fmt.Println(newTask.ToDescriptionString())
	},
}

func init() {
	rootCmd.AddCommand(chsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
