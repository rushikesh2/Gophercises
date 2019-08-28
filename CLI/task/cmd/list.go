package cmd

import (
	"fmt"

	"github.com/rushikesh2/GolangTraining/Gophercises/CLI/task/db"
	"github.com/spf13/cobra"
)

var Tasklist = db.AllTasks

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your incomplete tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := Tasklist()
		if err != nil {
			fmt.Println("Something went wrong:", err.Error())
			//os.Exit(1)
			return
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete!")
			return
		}
		fmt.Println("You have the following tasks:")
		for i, task := range tasks {
			fmt.Printf("%d. %s, Key=%d\n", i+1, task.Value, task.Key)

		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
