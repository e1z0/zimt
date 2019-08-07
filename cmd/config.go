package cmd

import (
	"github.com/spf13/cobra"

	"github.com/radiohive/zimt/pkg/config"
)

// configCmd implements the `config` command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Prints zimt configuration",
	Run: func(cmd *cobra.Command, args []string) {
		config.NewMqttConfig().Print()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
