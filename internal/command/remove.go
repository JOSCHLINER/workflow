package command

import (
	"workflow/internal/handler"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(removeCmd)

	removeCmd.PersistentFlags().Int("id", 0, "Id of task.")
}

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Command to remove a task.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		if id == 0 {
			cmd.Help()
			return
		}

		handler.RemoveTask(id)
	},
}
