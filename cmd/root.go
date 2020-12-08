package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/cheggaaa/pb"
	"github.com/simse/qc/internal/output"
	"github.com/simse/qc/internal/strategy"

	"github.com/simse/qc/internal/convert"

	"github.com/simse/qc/internal/source"
	"github.com/spf13/cobra"
)

// Input stores input extension flag
var Input string

// Output stores output extension flag
var Output string

// Recursive stores flag indicating whether scan should be recursive
var Recursive bool

// Directory stores input directory flag
var Directory string

var rootCmd = &cobra.Command{
	Use:   "qc",
	Short: "qc is a tool to convert file formats.",
	Long:  `qc (QuickConvert) can convert file formats quickly and easily.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Verify input and output format flags
		// Get array of input formats
		inputExtensions := strings.Split(Input, ",")

		for _, inputExtension := range inputExtensions {
			format, formatError := strategy.GetFormat(inputExtension)

			if formatError != nil {
				output.Error("Unknown input format: " + inputExtension)
				return
			}

			// Add format aliases
			// TODO: disable with --strict
			for _, formatAlias := range format.Aliases {
				inputExtensions = append(inputExtensions, formatAlias)
			}
		}

		// Check cutput format is known
		if !strategy.KnownExtension(Output) {
			output.Error("Unknown output format: " + Output)
			return
		}

		// Check output format can be encoded
		outputFormat, _ := strategy.GetFormat(Output)

		if !outputFormat.EncoderAvailable {
			output.Error("Cannot convert to format: " + Output + ", no encoder available")
			return
		}

		outputFormat, formatError := strategy.GetFormat(Output)
		if formatError != nil {
			panic(formatError)
		}

		// Check for runtime flags

		// Check for and verify format flags

		files, scanError := source.Scan(Directory, Recursive, inputExtensions)
		if scanError != nil {
			panic(scanError)
		}

		conversions := convert.PrepareAll(files, outputFormat)

		bar := pb.StartNew(len(conversions))
		var wg sync.WaitGroup

		for _, conversion := range conversions {
			wg.Add(1)

			go func(conversion convert.Conversion) {
				convert.Do(conversion)
				bar.Increment()
				wg.Done()
			}(conversion)
		}

		wg.Wait()
		bar.Finish()

		output.Success("Converted " + fmt.Sprint(len(conversions)) + " files")

		//conversionStats := convert.Stats{}

		//fmt.Println(conversionStats)

		// fmt.Print(strategy.KnownExtension(Input))

		//output.Warning("Converted 8 files")
	},
}

func init() {
	// Get current path
	currentPath, pathError := os.Getwd()
	if pathError != nil {
		log.Println(pathError)
	}

	rootCmd.Flags().StringVarP(&Input, "in", "i", "", "File extension(s) to convert from")
	rootCmd.Flags().StringVarP(&Output, "out", "o", "", "File extension to convert to")
	rootCmd.Flags().BoolVarP(&Recursive, "recursive", "R", false, "Whether folder scan should enter child directories")
	rootCmd.Flags().StringVarP(&Directory, "dir", "D", currentPath, "Input directory")
}

// Execute handles the root command, that is no subcommand
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
