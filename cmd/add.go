package cmd

import (
	"fmt"

	"github.com/itpourya/NotifyTask/internal/tasks"
	"github.com/spf13/cobra"
)

var addcommand = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a task description.")
			return
		}
		description := args[0]
		task := tasks.Task{}
		task.AddTask(description)
	},
}
