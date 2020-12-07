package tiff

import (
	"io"

	"github.com/simse/qc/internal/format"
	"golang.org/x/image/tiff"
)

// Info returns an Info struct about this format
func Info() format.Info {
	return format.Info{
		Extension: "tiff",
		Aliases:   []string{},
		HumanName: "tiff",
		Encoder:   Encode,
		// Decoder:   Decode,
	}
}

// Decode converts the PNG image to a generic image object for encoding
func Decode() {

}

// Encode converts a generic image object to a PNG file
func Encode(writer io.Writer, decodeObject format.DecodeOutput) format.EncodeOutput {
	writeError := tiff.Encode(writer, decodeObject.Image, nil)

	return format.EncodeOutput{
		Status: writeError == nil,
	}
}
