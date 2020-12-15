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

	// Check if result exists
	outputFile, fileOpenError := os.Open(conversion.OutputPath)
	if fileOpenError == nil {
		outputFile.Close()

		return ConversionResult{
			Skipped:    true,
			SkipReason: "result exists",
		}
	}

	// Create file reader
	fileSource, fileSourceError := os.Open(conversion.InputPath)
	if fileSourceError != nil {
		return ConversionResult{
			Error:   fileSourceError.Error(),
			Errored: true,
		}
	}
	defer fileSource.Close()

	// Create standard object
	standardObject, standardObjectError := conversion.InputFormat.Decoder(fileSource)

	if standardObjectError != nil {
		return ConversionResult{
			Error:   standardObjectError.Error(),
			Errored: true,
		}
	}

	// Create file writer
	fileDestination, fileDestinationError := os.Create(conversion.OutputPath)
	if fileDestinationError != nil {
		return ConversionResult{
			Error:   fileDestinationError.Error(),
			Errored: true,
		}
	}
	defer fileDestination.Close()

	// Write standard object to destination
	conversion.OutputFormat.Encoder(fileDestination, standardObject)

	return ConversionResult{
		Success: true,
	}
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
	extension := source.GetExtension(filePath, false)
	/*extensionLower := strings.ToLower(extension)
	wasUpperCase := false

	if extension != extensionLower {
		extension = extensionLower
		wasUpperCase = true
	}*/
	extensionInstances := strings.Count(filePath, extension)

	/*fmt.Println(extension)
	fmt.Println(extensionInstances)
	fmt.Println(replaceNth(filePath, extension, outputFormat.Extension, extensionInstances))*/

	return replaceNth(filePath, extension, outputFormat.Extension, extensionInstances)
}

// Replace the nth occurrence of old in s by new.
func replaceNth(s, old, new string, n int) string {
	i := 0
	for m := 1; m <= n; m++ {
		x := strings.Index(s[i:], old)
		if x < 0 {
			break
		}
		i += x
		if m == n {
			return s[:i] + new + s[i+len(old):]
		}
		i += len(old)
	}
	return s
}
