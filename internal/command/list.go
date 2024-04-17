package command

import (
	"workflow/internal/handler"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("active", "a", false, "List only the active tasks.")
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "used to list all tasks",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		onlyActive, _ := cmd.Flags().GetBool("active")

		if onlyActive {
			handler.ListActiveTasks()
		} else {
			handler.ListAllTasks()
		}

	},
}
