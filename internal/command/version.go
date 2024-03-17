package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the Version of workflow",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("workflow version 0.1")
	},
}
