package command

import (
	"workflow/internal/handler"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(stopCmd)

	stopCmd.PersistentFlags().IntP("id", "i", 0, "Id of task.")
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Command to stop a task.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		if id == 0 {
			cmd.Help()
			return
		}

		handler.StopTask(id)
	},
}
