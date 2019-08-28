package cmd

import (
	"fmt"
	"strings"

	"github.com/rushikesh2/GolangTraining/Gophercises/CLI/task/db"

	"github.com/spf13/cobra"
)

//var AddTask = db.CreateTask

//addCommand adds tasks to the list and stores it in Bolt DB
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add tasks to your todo list.",
	Run: func(cmd *cobra.Command, Args []string) {
		task := strings.Join(Args, " ") //taken as slice together as an argument
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong", err.Error())
			return
		}
		fmt.Printf("Added \"%s\" to your task list.\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
