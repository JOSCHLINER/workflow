package command

import (
	"workflow/internal/handler"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.PersistentFlags().Int("id", 0, "Id of task.")
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Command to start a task.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		if id == 0 {
			cmd.Help()
			return
		}

		handler.StartTask(id)
	},
}
