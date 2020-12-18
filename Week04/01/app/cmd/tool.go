package cmd

import (
	"github.com/spf13/cobra"
)

// toolCmd represents the tool command
var toolCmd = &cobra.Command{
	Use:   "tool",
	Short: "app tool",
	Long: ``,
}

var addUserCmd = &cobra.Command{
	Use:   "addUser",
	Short: "add app user",
	Long: ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}


var addGroupCmd = &cobra.Command{
	Use:   "addGroup",
	Short: "add app group",
	Long: ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(toolCmd)
	toolCmd.AddCommand(addUserCmd,addGroupCmd)
}
