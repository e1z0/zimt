package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/radiohive/zimt/pkg/api"
	"github.com/radiohive/zimt/pkg/mqtt"
)

// brokerCmd implements the `broker` command
var brokerCmd = &cobra.Command{
	Use:   "broker",
	Short: "Prints broker details",
	Run: func(cmd *cobra.Command, args []string) {
		client := mqtt.NewClient()
		broker := api.GetBrokerDetails(client)
		defer client.Disconnect(0)

		f := "%-18v%s\n"
		fmt.Printf(f, "Version:", broker.Version)
		fmt.Printf(f, "Uptime:", broker.FormatUptime())
		fmt.Printf(f, "MessagesSent:", broker.MessagesSent)
		fmt.Printf(f, "MessagesReceived:", broker.MessagesReceived)
		fmt.Printf(f, "ClientsConnected:", broker.ClientsConnected)
		fmt.Printf(f, "ClientsTotal:", broker.ClientsTotal)
	},
}

func init() {
	rootCmd.AddCommand(brokerCmd)
}
