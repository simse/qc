package cmd

import (
	"os"

	"github.com/equinox-io/equinox"
	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
	"github.com/simse/qc/internal/update"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates qc to the latest version",
	Long:  `Updates qc to the latest version`,
	Run: func(cmd *cobra.Command, args []string) {
		w := wow.New(os.Stdout, spin.Get(spin.Dots), " Updating qc...")
		w.Start()

		// Perform update
		updateResult := update.Do()

		if updateResult == equinox.NotAvailableErr {
			w.PersistWith(spin.Spinner{Frames: []string{"👍"}}, " Already on latest version")
		} else if updateResult == nil {
			w.PersistWith(spin.Spinner{Frames: []string{"👍"}}, " qc updated to latest version")
		} else {
			w.PersistWith(spin.Spinner{Frames: []string{"👍"}}, " Error while updating qc")
		}
	},
}
