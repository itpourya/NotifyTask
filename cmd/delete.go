package cmd

import (
	"fmt"
	"strconv"

	"github.com/itpourya/NotifyTask/internal/tasks"
	"github.com/spf13/cobra"
)

var deletecommand = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a task ID.")
			return
		}
		taskID, err := strconv.Atoi(args[0])
		task := tasks.Task{}
		if err != nil {
			fmt.Println("Please provide a task ID.")
			return
		}

		task.DeleteTask(taskID)
	},
}
