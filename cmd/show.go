package cmd

import (
	"fmt"

	"github.com/itpourya/NotifyTask/internal/tasks"
	"github.com/spf13/cobra"
)

var showcommand = &cobra.Command{
	Use:   "show",
	Short: "show tasks list",
	Run: func(cmd *cobra.Command, args []string) {
		task := tasks.Task{}
		fmt.Println(task.ShowTask())
	},
}
