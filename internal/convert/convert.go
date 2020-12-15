package convert

import (
	"os"
	"strings"

	"github.com/simse/qc/internal/strategy"

	"github.com/simse/qc/internal/format"
	"github.com/simse/qc/internal/source"
)

// Conversion stores all information needed to perform a conversion
type Conversion struct {
	InputPath    string
	OutputPath   string
	InputFormat  format.Info
	OutputFormat format.Info
}

// Do completes a conversion given input path, output path, input format and output format
func Do(conversion Conversion) ConversionResult {
	// Check if conversion is possible
	// if not, it's not an error
	// the conversion command has already been verified
	// mark as a skip instead of error
	canConvert, _ := strategy.VerifyConversion(conversion.InputFormat, conversion.OutputFormat)

	if !canConvert {
		return ConversionResult{
			Skipped:    true,
			SkipReason: "no strategy available",
		}
	}

	// Create file reader
	fileSource, fileSourceError := os.Open(conversion.InputPath)
	if fileSourceError != nil {
		panic(fileSourceError)
	}
	defer fileSource.Close()

	// fmt.Print(conversion.InputPath)

	// Create file writer
	fileDestination, fileDestinationError := os.Create(conversion.OutputPath)
	if fileDestinationError != nil {
		panic(fileDestinationError)
	}
	defer fileDestination.Close()

	// Create standard object
	standardObject := conversion.InputFormat.Decoder(fileSource)

	// Write standard object to destination
	conversion.OutputFormat.Encoder(fileDestination, standardObject)

	return ConversionResult{}
}

// PrepareAll prepares an array of files for conversion
func PrepareAll(files []source.File, outputFormat format.Info) []Conversion {
	conversions := []Conversion{}

	for _, file := range files {
		inputFormat, _, _ := strategy.GetFormat(file.Extension)

		conversions = append(conversions, Conversion{
			InputPath:    file.Path,
			OutputPath:   NewPath(file.Path, inputFormat, outputFormat),
			InputFormat:  inputFormat,
			OutputFormat: outputFormat,
		})
	}

	return conversions
}

// NewPath creates a new path for a file after conversion
func NewPath(filePath string, inputFormat format.Info, outputFormat format.Info) string {
	trimmedString := strings.TrimSuffix(filePath, inputFormat.Extension)

	// trimmedString will be unchanged if suffix wasn't there, try aliases instead
	for _, alias := range inputFormat.Aliases {
		if trimmedString != filePath {
			break
		}

		trimmedString = strings.TrimSuffix(filePath, alias)
	}

	return (trimmedString + outputFormat.Extension)
}
