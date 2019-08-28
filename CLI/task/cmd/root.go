package cmd

import (
	"github.com/spf13/cobra"
)

//RootCmd use task command as a base command
var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI task manager",
}
