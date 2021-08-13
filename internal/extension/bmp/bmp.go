package bmp

import (
	"image"
	"io"

	"github.com/simse/qc/internal/format"
	"golang.org/x/image/bmp"
)

// Info returns an Info struct about this format
func Info() format.Info {
	return format.Info{
		Extension:        "bmp",
		Aliases:          []string{"dib"},
		HumanName:        "bmp",
		Library:          "native",
		Encoder:          Encode,
		Decoder:          Decode,
		EncoderAvailable: true,
		DecoderAvailable: true,
	}
}

// Decode converts the BMP image to a generic image object for encoding
func Decode(reader io.Reader) (interface{}, error) {
	imageObject, decodeError := bmp.Decode(reader)

	if decodeError != nil {
		return nil, decodeError
	}

	return imageObject, nil
}

// Encode converts a generic image object to a PNG file
func Encode(writer io.Writer, decodeObject interface{}) (interface{}, error) {
	writeError := bmp.Encode(writer, decodeObject.(image.Image))

	return format.EncodeOutput{
		Status: writeError == nil,
	}, nil
}
