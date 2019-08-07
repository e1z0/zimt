package cmd

import (
	"fmt"

	"github.com/radiohive/zimt/pkg/buildmeta"
	"github.com/spf13/cobra"
)

// versionCmd implements the `version` command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints zimt version and build information",
	Run: func(cmd *cobra.Command, args []string) {
		f := "%-11v%s\n"
		fmt.Printf(f, "GitTag:", buildmeta.GitTag)
		fmt.Printf(f, "GitCommit:", buildmeta.GitCommit)
		fmt.Printf(f, "GitBranch:", buildmeta.GitBranch)
		fmt.Printf(f, "BuildDate:", buildmeta.BuildDate)
		fmt.Printf(f, "Platform:", buildmeta.Platform)
		fmt.Printf(f, "Compiler:", buildmeta.Compiler)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
