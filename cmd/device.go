package cmd

import (
	"github.com/spf13/cobra"
)

// deviceCmd implements the `device` command
var deviceCmd = &cobra.Command{
	Use:   "device",
	Short: "Device related commands family",
}

func init() {
	rootCmd.AddCommand(deviceCmd)
}
