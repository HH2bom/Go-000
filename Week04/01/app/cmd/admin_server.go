package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// adminserverCmd represents the adminserver command
var adminserverCmd = &cobra.Command{
	Use:   "adminserver",
	Short: "app admin-server",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("adminserver called")
	},
}

func init() {
	rootCmd.AddCommand(adminserverCmd)
	adminserverCmd.Flags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.app.yaml)")
}
