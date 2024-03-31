package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var configFile string

var rootCmd = &cobra.Command{
	Use:  "server",
	Long: "Restful API for systemd services",
	Run: func(cmd *cobra.Command, args []string) {
		runApplication()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(serverCmd)
	rootCmd.PersistentFlags().StringVarP(&configFile, "conf", "", "", "config file path")
}
