package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
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

		// If no input extensions where given assume wildcard
		if Input == "" {
			// fmt.Println("No input format given, assuming: *")
			inputExtensions[0] = "*"
		} else {
			for _, inputExtension := range inputExtensions {
				format, _, formatError := strategy.GetFormat(inputExtension)

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
		}

		// Check cutput format is known
		if len(args) == 0 && Output == "" {
			output.Error("No output format given, cannot convert to nothing")
			return
		}

		var outputFormatString string
		if Output != "" {
			outputFormatString = Output
		} else {
			outputFormatString = args[0]
		}

		if !strategy.KnownExtension(Output) && !strategy.KnownExtension(args[0]) {
			output.Error("Unknown output format: " + outputFormatString)
			return
		}

		// Check output format can be encoded
		outputFormat, _, _ := strategy.GetFormat(outputFormatString)

		if !outputFormat.EncoderAvailable {
			output.Error("Cannot convert to format: " + outputFormatString + ", no encoder available")
			return
		}

		outputFormat, _, formatError := strategy.GetFormat(outputFormatString)
		if formatError != nil {
			panic(formatError)
		}

		// Add output format as a negative filter in folder scan
		inputExtensions = append(inputExtensions, "-"+outputFormatString)

		// Check for runtime flags

		// Check for and verify format flags

		// Convert path to usable path (relative to absolute)
		Directory, _ = filepath.Abs(Directory)

		files, scanError := source.Scan(Directory, Recursive, inputExtensions)
		if scanError != nil {
			panic(scanError)
		}

		conversions := convert.PrepareAll(files, outputFormat)
		var conversionResults []convert.ConversionResult

		//bar := pb.StartNew(len(conversions))
		spinner := wow.New(os.Stdout, spin.Get(spin.Dots), " (0/"+fmt.Sprint(len(conversions))+") Processing and converting files")
		var wg sync.WaitGroup

		// Start convert timer
		conversionStart := time.Now()

		spinner.Start()
		for _, conversion := range conversions {
			wg.Add(1)

			go func(conversion convert.Conversion) {
				conversionResults = append(conversionResults, convert.Do(conversion))
				spinner.Text(" (" + fmt.Sprint(len(conversionResults)) + "/" + fmt.Sprint(len(conversions)) + ") Processing and converting files")
				wg.Done()
			}(conversion)
		}

		wg.Wait()
		spinner.Persist()
		fmt.Print("\n")

		// Stop conversion timer
		elapsed := time.Since(conversionStart)

		// Calculate conversion statistics
		skipped := 0
		succeeded := 0

		for _, conversionResult := range conversionResults {
			if conversionResult.Skipped {
				skipped++
			}

			if conversionResult.Success {
				succeeded++
			}
		}

		// Output conversion statistics
		if skipped > 0 {
			fmt.Println("Skipped " + fmt.Sprint(skipped) + " files")
		}
		output.Success("Converted " + fmt.Sprint(succeeded) + " files in " + fmt.Sprint(elapsed))
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
