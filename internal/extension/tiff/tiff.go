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
		Library:   "native",
		Encoder:   Encode,
		// Decoder:   Decode,
		EncoderAvailable: true,
		DecoderAvailable: true,
	}
}

// Decode converts the JPG image to a generic image object for encoding
func Decode(reader io.Reader) format.DecodeOutput {
	imageObject, _ := tiff.Decode(reader)

	return format.DecodeOutput{
		Image: imageObject,
	}
}

// Encode converts a generic image object to a PNG file
func Encode(writer io.Writer, decodeObject format.DecodeOutput) format.EncodeOutput {
	writeError := tiff.Encode(writer, decodeObject.Image, nil)

	return format.EncodeOutput{
		Status: writeError == nil,
	}
}
