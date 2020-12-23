package cmd

import (
	"fmt"

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
		/*w := wow.New(os.Stdout, spin.Get(spin.Dots), " Updating qc...")
		w.Start()

		// Perform update
		updateResult := update.Do()

		if updateResult == equinox.NotAvailableErr {
			w.PersistWith(spin.Spinner{Frames: []string{"👍"}}, " Already on latest version")
		} else if updateResult == nil {
			w.PersistWith(spin.Spinner{Frames: []string{"👍"}}, " qc updated to latest version")
		} else {
			w.PersistWith(spin.Spinner{Frames: []string{"👍"}}, " Error while updating qc")
		}*/

		fmt.Println("Sorry, this version does not support self-updating.")
	},
}
