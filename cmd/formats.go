package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/simse/qc/internal/strategy"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(formatsCommand)
}

var formatsCommand = &cobra.Command{
	Use:   "formats",
	Short: "Shows all formats supported by qc",
	Long:  `Lists all formats supported by qc, including group and so on.`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, group := range strategy.ConversionGroups {
			green := color.New(color.FgGreen)
			red := color.New(color.FgRed)
			cyan := color.New(color.FgCyan)

			whiteUnderline := color.New(color.FgWhite).Add(color.Underline)
			whiteUnderline.Print(group.FormatType)
			fmt.Println("")

			for _, format := range group.Formats {
				fmt.Print(format.HumanName)
				fmt.Print("> encoder: ")

				if format.EncoderAvailable {
					green.Print("yes")
				} else {
					red.Print("no")
				}

				fmt.Print(", decoder: ")
				if format.DecoderAvailable {
					green.Print("yes")
				} else {
					red.Print("no")
				}

				fmt.Print(", library: ")
				if format.Library == "native" {
					cyan.Print(format.Library)
				} else {
					fmt.Print(format.Library)
				}

				fmt.Print("\n")
			}

			if len(strategy.ConversionGroups) > 1 {
				fmt.Print("\n\n")
			}
		}
	},
}
