package command

import (
	"fmt"
	"workflow/internal/handler"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.PersistentFlags().StringP("name", "n", "", "Name of the task.")
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Command to create a task.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			fmt.Println("Please provide a name for the task!")
			return
		}
		handler.CreateTask(name)
	},
}
