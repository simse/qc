package tiff

import (
	"image"
	"io"

	"github.com/simse/qc/internal/format"
	"golang.org/x/image/tiff"
)

// Info returns an Info struct about this format
func Info() format.Info {
	return format.Info{
		Extension:        "tiff",
		Aliases:          []string{"tif"},
		HumanName:        "tiff",
		Library:          "native",
		Encoder:          Encode,
		Decoder:          Decode,
		EncoderAvailable: true,
		DecoderAvailable: true,
	}
}

// Decode converts the JPG image to a generic image object for encoding
func Decode(reader io.Reader) (interface{}, error) {
	imageObject, _ := tiff.Decode(reader)

	return imageObject, nil
}

// Encode converts a generic image object to a PNG file
func Encode(writer io.Writer, decodeObject interface{}) (interface{}, error) {
	writeError := tiff.Encode(writer, decodeObject.(image.Image), nil)

	return format.EncodeOutput{
		Status: writeError == nil,
	}, nil
}
