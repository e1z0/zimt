package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/radiohive/zimt/pkg/api"
	"github.com/radiohive/zimt/pkg/config"
	"github.com/radiohive/zimt/pkg/mqtt"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Prints zimt configuration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("==> Local zimt configuration:")
		config.NewMqttConfig().Print()

		fmt.Println()

		fmt.Println("==> Remote bridge configuration:")
		client := mqtt.NewClient()
		api.GetBridgeConfig(client).Print()
		client.Disconnect(50)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
