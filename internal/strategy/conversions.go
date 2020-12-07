package strategy

import (
	"errors"

	"github.com/simse/qc/internal/extension/bmp"
	"github.com/simse/qc/internal/extension/jpg"
	"github.com/simse/qc/internal/extension/png"
	"github.com/simse/qc/internal/extension/tiff"
	"github.com/simse/qc/internal/extension/webp"
	"github.com/simse/qc/internal/format"
)

// ConversionGroup represents a group of inter-convertible formats
type ConversionGroup struct {
	FormatType string
	Formats    []format.Info
}

// ConversionGroups contains all groups of formats that can be converted between
var ConversionGroups []ConversionGroup

func init() {
	// Create all the conversion groups
	ConversionGroups = []ConversionGroup{
		// Common image formats
		ConversionGroup{
			FormatType: "Images",
			Formats: []format.Info{
				jpg.Info(),
				png.Info(),
				webp.Info(),
				bmp.Info(),
				tiff.Info(),
			},
		},

		// TODO: Fonts
		// TODO: Common document formats
	}
}

// KnownExtension will check the conversion groups to see if an extension is known to qc
func KnownExtension(extension string) bool {
	_, formatError := GetFormat(extension)
	if formatError != nil {
		return false
	}

	return true
}

// GetFormat will return format info from extension string
func GetFormat(extension string) (format.Info, error) {
	for _, group := range ConversionGroups {
		for _, formatInfo := range group.Formats {
			if formatInfo.Extension == extension {
				return formatInfo, nil
			}

			for _, formatAlias := range formatInfo.Aliases {
				if formatAlias == extension {
					return formatInfo, nil
				}
			}
		}
	}

	return format.Info{}, errors.New("Format unknown")
}
