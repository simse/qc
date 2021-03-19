package cmd

import (
	"fmt"

	"github.com/simse/qc/internal/update"
	"github.com/spf13/cobra"
)

/*
// Version contains the qc version (filled by CI)
var Version string
*/
func init() {
	rootCmd.AddCommand(versionCmd)

	if update.Version == "" {
		update.Version = "dev"
	}
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the qc version",
	Long:  `The qc version number, indicating current version and whether the version is dirty.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(update.Version)
	},
}
