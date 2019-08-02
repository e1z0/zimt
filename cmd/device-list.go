package cmd

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	"github.com/radiohive/zimt/pkg/api"
	"github.com/radiohive/zimt/pkg/mqtt"
	"github.com/radiohive/zimt/pkg/structs"
)

// deviceListCmd implements the `device list` command
var deviceListCmd = &cobra.Command{
	Use:   "list",
	Short: "Prints all connected devices",
	Run: func(cmd *cobra.Command, args []string) {
		client := mqtt.NewClient()
		devices := api.GetBridgeConfigDevices(client)
		client.Disconnect(50)

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoFormatHeaders(false)
		titles, _ := structs.Titles(api.BridgeConfigDevice{}, "json", []string{"IEEEAddr", "FriendlyName", "Model", "ModelID", "ManufName", "SWBuildID", "DateCode"})
		table.SetHeader(titles)
		for _, d := range devices {
			table.Append([]string{d.IEEEAddr, d.FriendlyName, d.Model, d.ModelID, d.ManufName, d.SWBuildID, d.DateCode})
		}
		table.Render()
	},
}

func init() {
	deviceCmd.AddCommand(deviceListCmd)
}
