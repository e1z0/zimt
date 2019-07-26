package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/radiohive/zimt/pkg/config"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Prints zimt configuration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%#v\n", config.NewMqttConfig())
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
