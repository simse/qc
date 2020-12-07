package bmp

import (
	"io"

	"github.com/simse/qc/internal/format"
	"golang.org/x/image/bmp"
)

// Info returns an Info struct about this format
func Info() format.Info {
	return format.Info{
		Extension: "bmp",
		Aliases:   []string{},
		HumanName: "bmp",
		Encoder:   Encode,
		// Decoder:   Decode,
	}
}

// Decode converts the PNG image to a generic image object for encoding
func Decode() {

}

// Encode converts a generic image object to a PNG file
func Encode(writer io.Writer, decodeObject format.DecodeOutput) format.EncodeOutput {
	writeError := bmp.Encode(writer, decodeObject.Image)

	return format.EncodeOutput{
		Status: writeError == nil,
	}
}
