package cmd

import (
	"fmt"
	"strconv"

	"github.com/rushikesh2/Gophercises/CLI/todoz/store"

	"github.com/spf13/cobra"
)

var TempRemove = store.RemoveTasks

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "done is used to mark task as done.",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, id := range args {
			val, err := strconv.Atoi(id)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			ids = append(ids, val)
		}

		tasks, err := ShowAll()
		if err != nil {
			return
		}
		for _, id := range ids {
			if id < 0 || id > len(tasks) {
				fmt.Println("U have entered wrong id for deletion.")
				continue
			}
			task := tasks[id-1]
			//index out of bound
			err := TempRemove(task.Id)
			if err != nil {
				fmt.Printf("Task::%s cannot be done..:(", task.Task)
			}
			//print the task deleted
		}
	},
}

func init() {
	RootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
