package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/radiohive/zimt/pkg/api"
	"github.com/radiohive/zimt/pkg/mqtt"
)

// bridgeCmd implements the `bridge` command
var bridgeCmd = &cobra.Command{
	Use:   "bridge",
	Short: "Prints bridge version",
	Run: func(cmd *cobra.Command, args []string) {
		client := mqtt.NewClient()
		bridge := api.GetBridgeConfig(client)
		defer client.Disconnect(0)

		f := "%-13v%v\n"
		fmt.Printf(f, "Version:", bridge.Version)
		fmt.Printf(f, "Commit:", bridge.Commit)
		fmt.Printf(f, "Coordinator:", bridge.Coordinator)
		fmt.Printf(f, "LogLevel:", bridge.LogLevel)
		fmt.Printf(f, "PermitJoin:", bridge.PermitJoin)
	},
}

func init() {
	rootCmd.AddCommand(bridgeCmd)
}
