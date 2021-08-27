package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
	"github.com/simse/qc/internal/convert"
	"github.com/simse/qc/internal/output"
	"github.com/simse/qc/internal/strategy"

	"github.com/simse/qc/internal/source"
	"github.com/spf13/cobra"
)

var version string

// ConvertCommand is a generic struct containing all mandatory and optional inputs for a convert job
type ConvertCommand struct {
	Directory       string
	RecursiveScan   bool
	InputExtensions []string
	OutputExtension string
	Continue        bool
}

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
		// Set interrupt
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			sig := <-sigs
			fmt.Println()
			fmt.Println(sig)
			os.Exit(0)
		}()

		// Parse args and flags
		commandInput := HandleCommandInput(args)

		if !commandInput.Continue {
			return
		}

		/*
			// Print out standardised command
			fmt.Printf("qc --in %s --out %s --dir %s\n", commandInput.InputExtensions, commandInput.OutputExtension, commandInput.Directory)
		*/

		// Convert path to usable path (relative to absolute)
		Directory, _ = filepath.Abs(commandInput.Directory)

		// Prepare channel for streaming files scanned to conversion setup
		files := make(chan source.File)

		// In a separate go routine, scan the given directory
		go func() {
			scanError := source.Scan(Directory, commandInput.RecursiveScan, commandInput.InputExtensions, files)
			if scanError != nil {
				panic(scanError)
			}
		}()

		// Fetch output format
		outputFormat, group, _ := strategy.GetFormat(commandInput.OutputExtension)
		group.Setup()

		// Start convert timer
		conversionStart := time.Now()
		conversionResults := make(chan convert.ConversionResult)

		// Progress output setup, this is used by the spinner showing "(X/X)"
		conversionsStarted := 0
		resultsProcessed := 0

		// Ensure no more than X threads are running
		maxThreads := 24
		concurrentGoroutines := make(chan struct{}, maxThreads)
		var wg sync.WaitGroup

		// Set up file scan progress
		fileScanSpinner := wow.New(os.Stdout, spin.Get(spin.Line), " Found 0 files...")
		fileScanSpinner.Start()
		filesFound := 0

		// Start conversion for each file (received in channel from scanner)
		for file := range files {
			filesFound++

			fileScanSpinner.Text(" Found " + fmt.Sprint(filesFound) + " files...")

			wg.Add(1)
			go func(file source.File) {
				defer wg.Done()

				// Indicate a thread has started processing
				concurrentGoroutines <- struct{}{}
				conversionsStarted++

				// Set up and execute conversion job
				conversionJob := convert.Prepare(file, outputFormat)
				conversionResults <- convert.Do(conversionJob)

				// Indicate thread is done
				<-concurrentGoroutines
			}(file)
		}
		fileScanSpinner.Stop()
		output.Success("Found " + fmt.Sprint(filesFound) + " files")

		// When all conversions are done, close channel so program can display stats and exit
		go func() {
			wg.Wait()
			close(conversionResults)
		}()

		// Preapre spinner with default text
		spinner := wow.New(os.Stdout, spin.Get(spin.Line), " (0/"+fmt.Sprint(filesFound)+") Processing and converting files...")

		// Show spinner in terminal
		spinner.Start()

		// Calculate conversion statistics
		skipped := 0
		errored := 0
		warned := 0
		succeeded := 0

		// Handle every conversion result from channel
		for conversionResult := range conversionResults {
			resultsProcessed++

			spinner.Text(" (" + fmt.Sprint(resultsProcessed) + "/" + fmt.Sprint(filesFound) + ") Processing and converting files...")

			if conversionResult.Skipped {
				skipped++
			}

			if conversionResult.Success {
				succeeded++
			}

			if conversionResult.Warned {
				warned++
			}

			if conversionResult.Errored {
				errored++
			}
		}

		// spinner.Persist()
		fmt.Print("\n\n")

		// Stop conversion timer
		elapsed := time.Since(conversionStart)

		// Output conversion statistics
		if skipped > 0 {
			fmt.Println(fmt.Sprint(skipped) + " files skipped")
		}
		if warned > 0 {
			fmt.Println(fmt.Sprint(warned) + " conversions resulted in warnings")
		}
		if errored > 0 {
			fmt.Println(fmt.Sprint(errored) + " errors")
		}
		output.Success("Converted " + fmt.Sprint(succeeded) + " files in " + fmt.Sprint(elapsed))
	},
}

func init() {
	rootCmd.Flags().StringVarP(&Input, "in", "i", "", "File extension(s) to convert from")
	rootCmd.Flags().StringVarP(&Output, "out", "o", "", "File extension to convert to")
	rootCmd.Flags().BoolVarP(&Recursive, "recursive", "R", false, "Whether folder scan should enter child directories")
	rootCmd.Flags().StringVarP(&Directory, "dir", "D", "", "Input directory")
}

// Execute handles the root command, i.e. just `qc` with no provided subcommand
func Execute() {
	// Print version
	fmt.Printf("qc %s \n\n", version)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// HandleCommandInput will parse and verify command input
func HandleCommandInput(args []string) ConvertCommand {
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
				return ConvertCommand{
					Continue: false,
				}
			}

			// Add format aliases
			// TODO: disable with --strict
			inputExtensions = append(inputExtensions, format.Aliases...)
		}
	}

	// Check cutput format is known
	if len(args) == 0 && Output == "" {
		output.Error("No output format given, cannot convert to nothing")
		return ConvertCommand{
			Continue: false,
		}
	}

	var outputFormatString string
	if Output != "" {
		outputFormatString = Output
	} else {
		outputFormatString = args[0]
	}

	if !strategy.KnownExtension(Output) && !strategy.KnownExtension(args[0]) {
		output.Error("Unknown output format: " + outputFormatString)
		return ConvertCommand{
			Continue: false,
		}
	}

	// Check output format can be encoded
	outputFormat, _, _ := strategy.GetFormat(outputFormatString)

	if !outputFormat.EncoderAvailable {
		output.Error("Cannot convert to format: " + outputFormatString + ", no encoder available")
		return ConvertCommand{
			Continue: false,
		}
	}

	outputFormat, _, formatError := strategy.GetFormat(outputFormatString)
	if formatError != nil {
		panic(formatError)
	}

	// Find folder to look in
	if Directory == "" {
		// Get current path
		currentPath, pathError := os.Getwd()
		if pathError != nil {
			log.Println(pathError)
		}

		Directory = currentPath
	}

	return ConvertCommand{
		Directory:       Directory,
		InputExtensions: inputExtensions,
		OutputExtension: outputFormatString,
		RecursiveScan:   Recursive,
		Continue:        true,
	}
}
