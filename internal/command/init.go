package command

import (
	"workflow/internal/handler"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Command to initialize the program.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		handler.Initialize()
	},
}
