package server

import (
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

func init() {
	rootCmd.AddCommand(serverCmd)
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "", "", "config file path")
}

func Execute() error {
	return rootCmd.Execute()
}
