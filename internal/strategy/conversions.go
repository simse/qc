package strategy

import (
	"errors"

	"github.com/simse/qc/internal/format"
	"github.com/simse/qc/internal/strategy/codecs/jpg"
	"github.com/simse/qc/internal/strategy/codecs/png"
	"github.com/simse/qc/internal/strategy/images"
)

// ConversionGroup represents a group of inter-convertible formats
type ConversionGroup struct {
	FormatType string
	Formats    []format.Info
	Check      func() bool // Check for any dependencies before trying to use them
	Setup      func() bool // Setup things like libvips
}

// ConversionGroups contains all groups of formats that can be converted between
var ConversionGroups []ConversionGroup

func init() {
	// Create all the conversion groups
	ConversionGroups = []ConversionGroup{
		// Common image formats
		{
			FormatType: "Images",
			Formats: []format.Info{
				jpg.Info(),
				png.Info(),
			},
			Check: images.Check,
			Setup: images.Setup,
		},

		// TODO: Fonts
		// TODO: Common document formats
	}
}

// KnownExtension will check the conversion groups to see if an extension is known to qc
func KnownExtension(extension string) bool {
	_, _, formatError := GetFormat(extension)

	return formatError == nil
}

// GetFormat will return format info from extension string
func GetFormat(extension string) (format.Info, ConversionGroup, error) {
	for _, group := range ConversionGroups {
		for _, formatInfo := range group.Formats {
			if formatInfo.Extension == extension {
				return formatInfo, group, nil
			}

			for _, formatAlias := range formatInfo.Aliases {
				if formatAlias == extension {
					return formatInfo, group, nil
				}
			}
		}
	}

	return format.Info{}, ConversionGroup{}, errors.New("format unknown")
}

// VerifyConversion ensures input format can be converted to output format
func VerifyConversion(inputFormat format.Info, outputFormat format.Info) (bool, error) {
	_, inputFormatGroup, inputFormatError := GetFormat(inputFormat.Extension)
	if inputFormatError != nil {
		return false, errors.New("there was an error")
	}

	canConvert := false
	for _, format := range inputFormatGroup.Formats {
		if format.Extension == outputFormat.Extension {
			canConvert = true
		}

		for _, formatAlias := range format.Aliases {
			if formatAlias == outputFormat.Extension {
				canConvert = true
			}
		}
	}

	return canConvert, nil
}
