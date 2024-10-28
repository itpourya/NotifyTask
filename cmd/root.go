package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "taskcli",
	Short: "A CLI app to manage tasks and push notifications",
}

// Execute runs the root command
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.AddCommand(addcommand)
	rootCmd.AddCommand(deletecommand)
	rootCmd.AddCommand(showcommand)
	rootCmd.AddCommand(completecommand)
}
